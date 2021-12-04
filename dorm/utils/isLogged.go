package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"dorm/model"

	"github.com/garyburd/redigo/redis"
)

// IsLogged 判断用户是否已经登录 false 没有登录 true 已经登录
func IsLogged(r *http.Request) (ok bool, user *model.User) {
	//根据Cookie的name获取Cookie
	cookie, err := r.Cookie("jwt")
	if err != nil {
		return false, nil
	}
	token := cookie.Value

	// Redis 获取 user
	user = &model.User{}
	conn, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		fmt.Println("connect redis error:", err)
		return false, nil
	}
	defer conn.Close()
	userString, err := redis.String(conn.Do("GET", token))
	err = json.Unmarshal([]byte(userString), &user)
	// 未登录
	if err != nil {
		return false, nil
	}

	// 已登录
	return true, user
}
