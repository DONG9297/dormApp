package controller

import (
	"encoding/json"
	"net/http"

	"register/model"
	"register/utils"
)

func Register(w http.ResponseWriter, r *http.Request) {

	// json 解析请求
	/*
		请求格式
		{
			"phone": "18312345678",
			"stu_id": "1234567890",
			"name": "董",
			"gender": "女",
			"password": "e10adc3949ba59abbe56e057f20f883e"
		}
	*/
	user := model.User{}
	switch r.Method {
	case http.MethodPost:
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&user)
		if err != nil {
			result := utils.Result{
				Code:    http.StatusInternalServerError,
				Message: "解析请求失败",
				Data:    map[string]interface{}{"err": err},
			}
			result.Response(w)
			return
		}
	}

	// 判断手机号是否已注册
	if user.CheckPhone() {
		// 手机号存在
		result := utils.Result{
			Code:    http.StatusInternalServerError,
			Message: "手机号已注册",
		}
		result.Response(w)
		return
	}

	// 判断学号是否已注册
	if user.CheckStudentID() {
		// 学号存在
		result := utils.Result{
			Code:    http.StatusInternalServerError,
			Message: "学号已注册",
		}
		result.Response(w)
		return
	}

	//添加用户
	err := model.AddUser(&user)
	// 添加用户失败
	if err != nil {
		result := utils.Result{
			Code:    http.StatusInternalServerError,
			Message: "添加用户失败",
			Data:    map[string]interface{}{"err": err},
		}
		result.Response(w)
		return
	}
	// 添加用户成功
	result := utils.Result{
		Code:    http.StatusOK,
		Message: "成功",
		Data:    map[string]interface{}{"user": user},
	}
	result.Response(w)
}
