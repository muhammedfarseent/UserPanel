package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"WEAKS/testdose/model"
)

var Db *gorm.DB

func ConnectDB() {
	dsn := "root:new_password@tcp(127.0.0.1:3306)/USER?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect DB")
	}
	db.AutoMigrate(model.User{})
	Db = db
}
