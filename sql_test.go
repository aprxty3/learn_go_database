package main

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO customer (id, name) values ('budi', 'Budi')"

	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert data")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "select id, name from customer"

	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("id:", id, "name:", name)
	}
}

func TestQueryComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "select id, name, email, balance, rating,birth_date , married, created_at from customer"

	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birth_date sql.NullTime
		var created_at time.Time
		var married bool
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &married, &created_at)
		if err != nil {
			panic(err)
		}
		fmt.Println("id:", id, "name:", name, "email:", email, "balance:", balance, "rating:", rating, "birth_date:", birth_date, "married:", married, "created_at:", created_at)
	}
}

func TestQueryInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	username := "admin'; #"
	password := "salah"

	script := "select username from authentification where username= '" + username + "' and password='" + password + "' limit 1"

	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("sukses login:", username)
	} else {
		fmt.Println("gagal login")
	}
}

func TestQueryInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	username := "admin"
	password := "admin"

	script := "select username from authentification where username= ? and password=? limit 1"

	rows, err := db.QueryContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("sukses login:", username)
	} else {
		fmt.Println("gagal login")
	}
}

func TestExecSqlSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	username := "ajii"
	password := "jium"
	ctx := context.Background()
	script := "INSERT INTO authentification (username, password) values (?, ?)"

	_, err := db.ExecContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert data")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	email := "aji@gmail.com"
	comment := "wkwkwkwkwk begitu"
	ctx := context.Background()
	script := "INSERT INTO comment (email, comment) values (?, ?)"

	result, err := db.ExecContext(ctx, script, email, comment)
	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Your ID is", id)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "insert into comment (email, comment) values (?,?)"

	stmt, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := "aji" + strconv.Itoa(i) + "@gmail.com"
		comment := "wkwkwkwkwk begitu" + strconv.Itoa(i)

		result, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Your ID is", id)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	//do transaction
	script := "insert into comment (email, comment) values (?,?)"
	for i := 0; i < 10; i++ {
		email := "aji" + strconv.Itoa(i) + "@gmail.com"
		comment := "wkwkwkwkwk begitu" + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, script, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Your ID is", id)
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}
