package model

import (
	"math/rand"
	"strconv"
	"time"

	"register/utils"
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

func AddUser(user *User) (err error) {
	var uid int
	for {
		rand.Seed(time.Now().Unix())
		uid = rand.Intn(1e8) //1e8
		var temp int
		row := utils.Db.QueryRow("Select uid from users where uid=?", uid)
		err := row.Scan(&temp)
		if err != nil {
			break
		}
	}
	sqlStr := "insert into users(uid, phone, stu_id, name, gender, password) values(?,?,?,?,?,?)"
	_, err = utils.Db.Exec(sqlStr, strconv.Itoa(uid), user.Phone, user.StudentID, user.Name, user.Gender, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (user *User) CheckPhone() bool {
	sqlStr := "select user_id, uid, phone, stu_id, name, gender, password from users where phone = ?"
	row := utils.Db.QueryRow(sqlStr, user.Phone)
	err := row.Scan(&user.ID, &user.UID, &user.Phone, &user.StudentID, &user.Name, &user.Gender, &user.Password)
	// 手机号未注册
	if err != nil {
		return false
	}
	return true
}

func (user *User) CheckStudentID() bool {
	sqlStr := "select user_id, uid, phone, stu_id, name, gender, password from users where stu_id = ?"
	row := utils.Db.QueryRow(sqlStr, user.StudentID)
	err := row.Scan(&user.ID, &user.UID, &user.Phone, &user.StudentID, &user.Name, &user.Gender, &user.Password)
	// 学号未注册
	if err != nil {
		return false
	}
	return true
}
