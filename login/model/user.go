package model

import (
	"login/utils"
)

// User 用户格式
type User struct {
	ID        int    `json:"id"`
	UID       string `json:"uid"`
	Phone     string `json:"phone"`
	StudentID string `json:"stu_id"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	Password  string `json:"password"`
}

func GetUserByPhone(phone string) (user *User, err error) {
	sqlStr := "select user_id, uid, phone, stu_id, name, gender, password from users where phone = ?"
	row := utils.Db.QueryRow(sqlStr, phone)
	user = &User{}
	err = row.Scan(&user.ID, &user.UID, &user.Phone, &user.StudentID, &user.Name, &user.Gender, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByID(ID int) (user *User, err error) {
	sqlStr := "select user_id, uid, phone, stu_id, name, gender, password from users where user_id = ?"
	row := utils.Db.QueryRow(sqlStr, ID)
	user = &User{}
	err = row.Scan(&user.ID, &user.UID, &user.Phone, &user.StudentID, &user.Name, &user.Gender, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
