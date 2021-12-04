package controller

import (
	"dorm/dao"
	"net/http"

	"dorm/utils"
)

func GetDormList(w http.ResponseWriter, r *http.Request) {
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

	// 获取宿舍信息 宿舍楼名 空余床位数1，2，3，4... 数量
	dormList, err := dao.GetDormList(user.Gender)
	if err != nil {
		result := utils.Result{
			Code:    http.StatusInternalServerError,
			Message: "获取宿舍信息失败",
		}
		result.Response(w)
		return
	}
	buildingList, err := dao.GetBuildingList(user.Gender)
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
			"dorm_list":     dormList,
			"building_list": buildingList,
		},
	}
	result.Response(w)
	return

}
