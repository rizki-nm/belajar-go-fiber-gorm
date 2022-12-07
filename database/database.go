package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/belajar_go_fiber?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicf("Cannot connect to database")
	}
	log.Println("Connected to database")
}
