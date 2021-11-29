package utils

// Session 结构
type Session struct {
	SessionID string `json:"session_id"`
	UserID    int    `json:"user_id"`
}

//AddSession 向数据库中添加 Session
func AddSession(sess *Session) error {
	sqlStr := "insert into sessions (session_id, user_id) values(?,?)"
	_, err := Db.Exec(sqlStr, sess.SessionID, sess.UserID)
	if err != nil {
		return err
	}
	return nil
}

//DeleteSession 删除数据库中的 Session
func DeleteSession(sessionID string) error {
	sqlStr := "delete from sessions where session_id = ?"
	_, err := Db.Exec(sqlStr, sessionID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteSessionByUser 删除数据库中对应 UserID 的 Session
func DeleteSessionByUser(userID int) error {
	sqlStr := "delete from sessions where user_id = ?"
	_, err := Db.Exec(sqlStr, userID)
	if err != nil {
		return err
	}
	return nil
}

//GetSession 根据session的Id值从数据库中查询 Session
func GetSession(sessionID string) (sess *Session) {
	sqlStr := "select session_id, user_id from sessions where session_id = ?"
	row := Db.QueryRow(sqlStr, sessionID)
	//扫描数据库中的字段值为Session的字段赋值
	sess = &Session{}
	err := row.Scan(&sess.SessionID, &sess.UserID)
	if err != nil {
		return nil
	}
	return sess
}
