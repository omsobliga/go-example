package model

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db2 *gorm.DB

func initDB2() {
	dsn := "root:root@tcp(localhost:3306)/web?charset=utf8mb4&parseTime=True"
	var err error
	db2, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func DB2() *gorm.DB {
	return db2
}

func init() {
	initDB2()
	log.Printf("init db success\n")
}
