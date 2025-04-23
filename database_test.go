package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/mysql_go")
	if err != nil {
		panic(err)
	}

	/*
		gunakan DB
	*/

	defer db.Close()
}
