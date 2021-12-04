package controller

import (
	"crypto/rand"
	"dorm/dao"
	"dorm/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"dorm/utils"
)

func ChooseDorm(w http.ResponseWriter, r *http.Request) {
	// 判断是否登录
	ok, user := utils.IsLogged(r)
	if !ok {
		result := utils.Result{
			Code:    http.StatusUnauthorized,
			Message: "未登录",
		}
		result.Response(w)
		return
	}
	// 解析请求
	var data struct {
		BuildingName string `json:"building"`
		StuCode0     string `json:"stucode0"`
		StuCode1     string `json:"stucode1"`
		StuCode2     string `json:"stucode2"`
	}
	switch r.Method {
	case http.MethodPost:
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&data)
		if err != nil {
			result := utils.Result{
				Code:    http.StatusInternalServerError,
				Message: "解析请求失败",
			}
			result.Response(w)
			return
		}
	}
	var stuCodes = [3]string{data.StuCode0, data.StuCode1, data.StuCode2}

	// 判断参数是否合法
	//temp, err := dao.GetUserByID(user.ID)
	//if err != nil || temp.ID != user.ID {
	//	result := utils.Result{
	//		Code:    http.StatusInternalServerError,
	//		Message: "请求不合法",
	//	}
	//	result.Response(w)
	//	return
	//}

	building, err := dao.GetBuildingByName(data.BuildingName)
	if err != nil || building.Name != data.BuildingName {
		result := utils.Result{
			Code:    http.StatusInternalServerError,
			Message: "宿舍楼错误",
		}
		result.Response(w)
		return
	}

	users := make(map[int]*model.User)
	users[user.ID] = user
	for _, stucode := range stuCodes {
		if stucode != "" {
			coUser, err := dao.GetUserByUID(stucode)
			if err != nil || coUser.StudentID != stucode {
				result := utils.Result{
					Code:    http.StatusInternalServerError,
					Message: "同住人不存在",
					Data:    map[string]interface{}{"user": coUser},
				}
				result.Response(w)
				return
			}
			if coUser.Gender != user.Gender {
				result := utils.Result{
					Code:    http.StatusInternalServerError,
					Message: "同住人性别错误",
					Data:    map[string]interface{}{"user": coUser},
				}
				result.Response(w)
				return
			}
			if dao.HasUserChosenDorm(coUser.ID) {
				result := utils.Result{
					Code:    http.StatusInternalServerError,
					Message: "同住人已选宿舍",
					Data:    map[string]interface{}{"user": coUser},
				}
				result.Response(w)
				return
			}
			users[coUser.ID] = coUser
		}
	}

	// 生成订单
	// 创建生成订单的时间
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	// 生成UUID作为订单的id
	u := new([16]byte)
	_, err = rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}
	u[8] = (u[8] | 0x40) & 0x7F
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])

	// 添加订单
	order := &model.Order{
		ID:         uuid,
		UserID:     user.ID,
		Count:      len(users),
		BuildingID: building.ID,
		Gender:     user.Gender,
		CreateTime: timeStr,
		State:      0,
	}
	err = dao.AddOrder(order)
	if err != nil {
		result := utils.Result{
			Code:    http.StatusInternalServerError,
			Message: "生成订单失败",
		}
		result.Response(w)
		return
	}

	// 添加订单项
	for _, user := range users {
		orderItem := &model.OrderItem{
			OrderID: uuid,
			UserID:  user.ID,
		}
		err = dao.AddOrderItem(orderItem)
		if err != nil {
			result := utils.Result{
				Code:    http.StatusInternalServerError,
				Message: "生成订单失败",
			}
			result.Response(w)
			return
		}
	}

	//	返回结果
	result := utils.Result{
		Code:    http.StatusOK,
		Message: "成功",
	}
	result.Response(w)
	return
}
