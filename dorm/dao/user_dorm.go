package dao

import "dorm/utils"

func AddUserIntoDorm(userID, dormID int) (err error) {
	sqlStr := "insert into user_dorm(user_id, dorm_id) values(?,?)"
	_, err = utils.Db.Exec(sqlStr, userID, dormID)
	return err
}

func GetDormIDByUserID(userID int) (dormID int, err error) {
	sqlStr := "select dorm_id from user_dorm where user_id = ?"
	row := utils.Db.QueryRow(sqlStr, userID)
	err = row.Scan(&dormID)
	if err != nil {
		return dormID, err
	}
	return dormID, nil
}

func HasUserChosenDorm(userID int) bool {
	_, err := GetDormIDByUserID(userID)
	if err != nil {
		return false
	}
	return true
}

func GetUserIDsByDormID(dormID int) (dormIDs []int, err error) {
	sql := "select user_id from user_dorm where dorm_id = ?"
	rows, err := utils.Db.Query(sql, dormID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var dormID int
		err := rows.Scan(&dormID)
		if err != nil {
			return nil, err
		}
		dormIDs = append(dormIDs, dormID)
	}
	return dormIDs, nil
}
