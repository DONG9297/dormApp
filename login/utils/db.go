package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	// todo mysql
	Db, err = sql.Open("mysql",
		"root:123456@tcp(db:3336)/userdb")
	if err != nil {
		panic(err.Error())
	}
}
