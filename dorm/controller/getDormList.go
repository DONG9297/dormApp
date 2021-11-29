package controller

import (
	"encoding/json"
	"net/http"

	"dorm/model"
	"dorm/utils"
)

func GetDormList(w http.ResponseWriter, r *http.Request) {
	// 解析请求
	var data struct {
		Gender string `json:"gender"`
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

	// 获取宿舍信息 宿舍楼名 空余床位数1，2，3，4... 数量
	dormList, err := model.GetDormList(data.Gender)
	if err != nil {
		result := utils.Result{
			Code:    http.StatusInternalServerError,
			Message: "获取宿舍信息失败",
		}
		result.Response(w)
		return
	}
	// 返回结果
	result := utils.Result{
		Code:    http.StatusOK,
		Message: "成功",
		Data: map[string]interface{}{
			"dorm_list": dormList,
		},
	}
	result.Response(w)
	return

}
