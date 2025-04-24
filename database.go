package main

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	target := "root:Awanmendung12?@tcp(localhost:3306)/mysql_go?parseTime=true"
	db, err := sql.Open("mysql", target)

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
