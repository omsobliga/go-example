package model

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() (err error) {
	dsn := "root:root@tcp(localhost:3306)/web?charset=utf8mb4&parseTime=True"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return nil
}

func DB() *sql.DB {
	return db
}

func init() {
	initDB()
	log.Printf("init db success\n")
}
