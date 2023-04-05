package dao

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

var db *sql.DB

const dsn = "root:123456@tcp(localhost:4000)/student?charset=utf8mb4&parseTime=True&loc=Local"

func init() {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logrus.Error("[mysql][init] open error: ", err)
		return
	}
	if err = db.Ping(); err != nil {
		logrus.Error("[mysql][init] ping error: ", err)
		return
	}
}

func Create() error {
	sqlStr := ``
	exec, err := db.Exec(sqlStr, "")
	if err != nil {
		return err
	}
	exec.RowsAffected()

	return nil
}
