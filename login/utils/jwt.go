package utils

import (
	"errors"
	"time"

	"login/model"

	"github.com/dgrijalva/jwt-go"
)

const (
	SecretKey = "Dong Rui"
)

// GenerateToken 生成 token
func GenerateToken(user *model.User) (string, error) {
	// 生成token
	claim := jwt.MapClaims{
		"exp":    time.Now().Add(time.Hour * time.Duration(24)).Unix(),
		"iat":    time.Now().Unix(),
		"id":     user.ID,
		"uid":    user.UID,
		"phone":  user.Phone,
		"stu_id": user.StudentID,
		"name":   user.Name,
		"gender": user.Gender,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	//加密
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken 解析 token
func ParseToken(tokenString string) (user *model.User, err error) {
	user = &model.User{}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("cannot convert claim to mapclaim")
		return
	}
	//验证token，如果token被修改过则为false
	if !token.Valid {
		err = errors.New("token is invalid")
		return
	}

	user.ID = int(claim["id"].(float64))
	user.UID = claim["uid"].(string)
	user.Phone = claim["phone"].(string)
	user.StudentID = claim["stu_id"].(string)
	user.Name = claim["name"].(string)
	user.Gender = claim["gender"].(string)

	return
}
