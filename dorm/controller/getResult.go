package controller

import (
	"dorm/dao"
	"net/http"

	"dorm/utils"
)

func GetResult(w http.ResponseWriter, r *http.Request) {
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

	dormID, err := dao.GetDormIDByUserID(user.ID)
	if err != nil {
		result := utils.Result{
			Code:    http.StatusInternalServerError,
			Message: "未查询到宿舍",
		}
		result.Response(w)
		return
	}
	dormInfo, err := dao.GetDormInfoByDormID(dormID)
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
			"user":      user,
			"dorm_info": dormInfo,
		},
	}
	result.Response(w)
	return
}

// ProcessOrder 处理时间早于timeStr发生的订单
func ProcessOrder(timeStr string) {
	// 获取待处理订单
	orders, err := dao.GetUnprocessedOrdersBefore(timeStr)
	if err != nil {
		return
	}
	for _, order := range orders {
		building, err := dao.GetBuildingByID(order.BuildingID)
		if err != nil {
			dao.UpdateOrderState(order.ID, 2)
			return
		}
		// 订单中用户是否都未选宿舍
		flag := true
		items, err := dao.GetItemsByOrderID(order.ID)
		if err != nil {
			dao.UpdateOrderState(order.ID, 2)
			return
		}
		for _, item := range items {
			if dao.HasUserChosenDorm(item.UserID) {
				dao.UpdateOrderState(order.ID, 2)
				flag = false
				break
			}
		}
		if !flag {
			continue
		}
		// 获取满足订单条件的宿舍列表
		dormInfos, err := dao.GetAvailableDormInfos(order.Count, building.Name, order.Gender)
		if err != nil {
			dao.UpdateOrderState(order.ID, 2)
			return
		}
		// 如果宿舍列表不为空
		if len(dormInfos) > 0 {
			dormInfo := dormInfos[0]
			// 将选宿舍信息加入学生宿舍表
			items, _ := dao.GetItemsByOrderID(order.ID)
			for _, item := range items {
				dao.AddUserIntoDorm(item.UserID, dormInfo.DormID)
			}
			// 更新宿舍空床数
			availableBeds := dormInfo.AvailableBeds - order.Count
			dao.UpdateDormAvailableBeds(dormInfo.DormID, availableBeds)
			// 更新订单状态
			dao.UpdateOrderState(order.ID, 1)
		} else {
			dao.UpdateOrderState(order.ID, 2)
		}
	}
}
