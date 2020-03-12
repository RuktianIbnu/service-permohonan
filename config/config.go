package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func DbConn() *sql.DB{
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/dagun?charset=utf8&parseTime=True&loc=Local")
	if err  != nil {
		panic("failed to connect to database")
	}

	return db
}