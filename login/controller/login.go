package controller

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"login/model"
	"login/utils"
)

func IsLogged(w http.ResponseWriter, r *http.Request) {

	// json 解析请求
	var sessionID string
	switch r.Method {
	case http.MethodPost:
		session := utils.Session{}
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&session)
		if err != nil {
			result := utils.Result{
				Code:    http.StatusInternalServerError,
				Message: "解析请求失败",
			}
			result.Response(w)
			return
		}
		sessionID = session.SessionID
	}

	session := utils.GetSession(sessionID)
	if session != nil && session.SessionID == sessionID {
		user := model.GetUserByID(session.UserID)
		if user != nil && user.ID == session.UserID {
			// 返回已登录消息
			result := utils.Result{
				Code:    http.StatusOK,
				Message: "已登录",
				Data:    map[string]interface{}{"user": user},
			}
			result.Response(w)
			return
		}
	}

	// 返回未登录消息
	result := utils.Result{
		Code:    http.StatusInternalServerError,
		Message: "未登录",
	}
	result.Response(w)
	return

}

func Login(w http.ResponseWriter, r *http.Request) {

	// json 解析请求
	var phone, password string
	switch r.Method {
	case http.MethodPost:
		user := model.User{}
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&user)
		if err != nil {
			result := utils.Result{
				Code:    http.StatusInternalServerError,
				Message: "解析请求失败",
			}
			result.Response(w)
			return
		}
		phone = user.Phone
		password = user.Password
	}

	// 判断用户名和密码是否正确
	user := model.GetUserByPhone(phone)
	if user == nil || user.Password != password {
		//手机号或密码不正确
		result := utils.Result{
			Code:    http.StatusInternalServerError,
			Message: "手机号或密码不正确",
		}
		result.Response(w)
		return
	}

	// Session
	// 生成UUID作为Session的id
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}
	u[8] = (u[8] | 0x40) & 0x7F
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])

	// 将Session保存到数据库中
	session := utils.Session{
		SessionID: uuid,
		UserID:    user.ID,
	}
	utils.AddSession(&session)

	// 返回SessionID和成功消息
	result := utils.Result{
		Code:    200,
		Message: "登陆成功",
		Data: map[string]interface{}{
			"user":    user,
			"session": session,
		},
	}
	result.Response(w)
	return
}

func Logout(w http.ResponseWriter, r *http.Request) {

	// json 解析请求
	var sessionID string
	switch r.Method {
	case http.MethodPost:
		session := utils.Session{}
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&session)
		if err != nil {
			result := utils.Result{
				Code:    http.StatusInternalServerError,
				Message: "解析请求失败",
			}
			result.Response(w)
			return
		}
		sessionID = session.SessionID
	}

	// 删除数据库中与之对应的Session
	session := utils.GetSession(sessionID)
	if session == nil {
		result := utils.Result{
			Code:    http.StatusInternalServerError,
			Message: "未登录",
		}
		result.Response(w)
		return
	}
	err := utils.DeleteSessionByUser(session.UserID)
	if err != nil {
		result := utils.Result{
			Code:    http.StatusInternalServerError,
			Message: "未登录",
		}
		result.Response(w)
		return
	}

	// 返回成功消息
	result := utils.Result{
		Code:    http.StatusOK,
		Message: "退出成功",
	}
	result.Response(w)
	return
}
