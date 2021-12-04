package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"login/dao"
	"login/model"
	"login/utils"

	"github.com/garyburd/redigo/redis"
)

func IsLogged(w http.ResponseWriter, r *http.Request) {
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
	// 已登录
	result := utils.Result{
		Code:    http.StatusOK,
		Message: "已登录",
		Data:    map[string]interface{}{"user": user},
	}
	result.Response(w)
	return
}

func Login(w http.ResponseWriter, r *http.Request) {

	// json 解析请求
	/*
		请求格式
		{
			"phone": "18312345678",
			"password": "e10adc3949ba59abbe56e057f20f883e"
		}
	*/
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
				Data:    map[string]interface{}{"err": err},
			}
			result.Response(w)
			return
		}
		phone = user.Phone
		password = user.Password
	}

	// 判断用户名和密码是否正确
	user, err := dao.GetUserByPhone(phone)
	if err != nil || user.Password != password {
		//手机号或密码不正确
		result := utils.Result{
			Code:    http.StatusInternalServerError,
			Message: "手机号或密码不正确",
		}
		result.Response(w)
		return
	}

	// token
	token, err := utils.GenerateToken(user)
	if err != nil {
		result := utils.Result{
			Code:    http.StatusInternalServerError,
			Message: "生成token失败",
		}
		result.Response(w)
		return
	}
	userJson, _ := json.Marshal(user)
	conn, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		fmt.Println("connect redis error:", err)
		return
	}
	defer conn.Close()

	_, err = conn.Do("SET", token, userJson)
	if err != nil {
		fmt.Println("redis set error:", err)
	}

	// 返回 token 和成功消息
	result := utils.Result{
		Code:    200,
		Message: "登陆成功",
		Data: map[string]interface{}{
			"token": token,
		},
	}
	result.Response(w)
	return

	//// Session
	//// 生成UUID作为Session的id
	//u := new([16]byte)
	//_, err = rand.Read(u[:])
	//if err != nil {
	//	result := utils.Result{
	//		Code:    http.StatusInternalServerError,
	//		Message: "无法登录",
	//		Data:    map[string]interface{}{"err": err},
	//	}
	//	result.Response(w)
	//	return
	//}
	//u[8] = (u[8] | 0x40) & 0x7F
	//u[6] = (u[6] & 0xF) | (0x4 << 4)
	//uuid := fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	//
	//// 将Session保存到数据库中
	//session := utils.Session{
	//	SessionID: uuid,
	//	UserID:    user.ID,
	//}
	//utils.AddSession(&session)
	//
	//// 返回SessionID和成功消息
	//result := utils.Result{
	//	Code:    200,
	//	Message: "登陆成功",
	//	Data: map[string]interface{}{
	//		"user":    user,
	//		"session": session,
	//	},
	//}
	//result.Response(w)
	//return
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("jwt")
	token := cookie.Value
	// 判断是否登录
	ok, _ := utils.IsLogged(r)
	if !ok {
		result := utils.Result{
			Code:    http.StatusUnauthorized,
			Message: "未登录",
		}
		result.Response(w)
		return
	}

	// 删除 Redis 中的数据
	conn, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		fmt.Println("connect redis error:", err)
		return
	}
	defer conn.Close()
	_, err = conn.Do("DEL", token)
	if err != nil {
		fmt.Println("redis del error:", err)
	}

	// 返回成功消息
	result := utils.Result{
		Code:    http.StatusOK,
		Message: "退出成功",
	}
	result.Response(w)
	return
}

func CheckPassword(w http.ResponseWriter, r *http.Request) {

	// json 解析请求
	/*
		请求格式
		{
			"phone": "18312345678",
			"password": "e10adc3949ba59abbe56e057f20f883e"
		}
	*/
	var data struct {
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}
	switch r.Method {
	case http.MethodPost:
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&data)
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

	// 判断用户名和密码是否正确
	user, err := dao.GetUserByPhone(data.Phone)
	if err != nil || user.Password != data.Password {
		//手机号或密码不正确
		result := utils.Result{
			Code:    http.StatusInternalServerError,
			Message: "手机号或密码不正确",
		}
		result.Response(w)
		return
	}

	// 返回成功消息
	result := utils.Result{
		Code:    http.StatusOK,
		Message: "正确",
	}
	result.Response(w)
	return
}
