package dao

import (
	"dorm/model"
	"dorm/utils"
)

func GetUserByPhone(phone string) (user *model.User, err error) {
	sqlStr := "select user_id, uid, phone, stu_id, name, gender, password from users where phone = ?"
	row := utils.Db.QueryRow(sqlStr, phone)
	user = &model.User{}
	err = row.Scan(&user.ID, &user.UID, &user.Phone, &user.StudentID, &user.Name, &user.Gender, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByID(ID int) (user *model.User, err error) {
	sqlStr := "select user_id, uid, phone, stu_id, name, gender, password from users where user_id = ?"
	row := utils.Db.QueryRow(sqlStr, ID)
	user = &model.User{}
	err = row.Scan(&user.ID, &user.UID, &user.Phone, &user.StudentID, &user.Name, &user.Gender, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByUID(UID string) (user *model.User, err error) {
	sqlStr := "select user_id, uid, phone, stu_id, name, gender, password from users where uid = ?"
	row := utils.Db.QueryRow(sqlStr, UID)
	user = &model.User{}
	err = row.Scan(&user.ID, &user.UID, &user.Phone, &user.StudentID, &user.Name, &user.Gender, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
