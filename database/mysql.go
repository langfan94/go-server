package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func InitMysql() {
	var err error
	Db, err = sql.Open("mysql", "root:x804311655@/database1")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	err = Db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	Db.SetMaxIdleConns(20)
	Db.SetMaxOpenConns(20)
}
