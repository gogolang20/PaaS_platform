package dao

import (
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"
)

type IMysqlDB interface {
	InsertOne(ctx context.Context) error
}

type MysqlDB struct {
	db *sql.DB
}

func NewMysqlDB() IMysqlDB {
	db, err := sql.Open("mysql", MysqlDSN)
	if err != nil {
		logrus.Error("[mysql][init] open error: ", err)
		return nil
	}
	if err = db.Ping(); err != nil {
		logrus.Error("[mysql][init] ping error: ", err)
		return nil
	}

	return &MysqlDB{db: db}
}

func (mdb *MysqlDB) InsertOne(ctx context.Context) error {

	result, err := mdb.db.Exec("")
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {

	}

	return nil
}

var db *sql.DB

const (
	MysqlDSN = "root:123456@tcp(localhost:4000)/student?charset=utf8mb4&parseTime=True&loc=Local"
)

func init() {
	db, err := sql.Open("mysql", MysqlDSN)
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
