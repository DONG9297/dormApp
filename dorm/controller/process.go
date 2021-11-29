package controller

import (
	"encoding/json"
	"net/http"

	"dorm/model"
	"dorm/utils"
)

func GetResult(w http.ResponseWriter, r *http.Request) {
	// 解析请求
	var data struct {
		UserID int `json:"user_id"`
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

	dormID, err := model.GetDormIDByUserID(data.UserID)
	if err != nil {
		result := utils.Result{
			Code:    http.StatusInternalServerError,
			Message: "未查询到宿舍",
		}
		result.Response(w)
		return
	}
	dormInfo, err := model.GetDormInfoByDormID(dormID)
	if err != nil {
		result := utils.Result{
			Code:    http.StatusInternalServerError,
			Message: "未查询到宿舍",
		}
		result.Response(w)
		return
	}

	result := utils.Result{
		Code:    http.StatusOK,
		Message: "成功",
		Data: map[string]interface{}{
			"dorm_info": dormInfo,
		},
	}
	result.Response(w)
	return
}

// ProcessOrder 处理时间早于timeStr发生的订单
func ProcessOrder(timeStr string) {
	// 获取待处理订单
	orders, err := model.GetUnprocessedOrdersBefore(timeStr)
	if err != nil {
		return
	}
	for _, order := range orders {
		building, err := model.GetBuildingByID(order.BuildingID)
		if err != nil {
			model.UpdateOrderState(order.ID, 2)
			return
		}
		// 订单中用户是否都未选宿舍
		flag := true
		items, err := model.GetItemsByOrderID(order.ID)
		if err != nil {
			model.UpdateOrderState(order.ID, 2)
			return
		}
		for _, item := range items {
			if model.HasUserChosenDorm(item.UserID) {
				model.UpdateOrderState(order.ID, 2)
				flag = false
				break
			}
		}
		if !flag {
			continue
		}
		// 获取满足订单条件的宿舍列表
		dormInfos, err := model.GetAvailableDormInfos(order.Count, building.Name, order.Gender)
		if err != nil {
			model.UpdateOrderState(order.ID, 2)
			return
		}
		// 如果宿舍列表不为空
		if len(dormInfos) > 0 {
			dormInfo := dormInfos[0]
			// 将选宿舍信息加入学生宿舍表
			items, _ := model.GetItemsByOrderID(order.ID)
			for _, item := range items {
				model.AddUserIntoDorm(item.UserID, dormInfo.DormID)
			}
			// 更新宿舍空床数
			availableBeds := dormInfo.AvailableBeds - order.Count
			model.UpdateDormAvailableBeds(dormInfo.DormID, availableBeds)
			// 更新订单状态
			model.UpdateOrderState(order.ID, 1)
		} else {
			model.UpdateOrderState(order.ID, 2)
		}
	}
}
