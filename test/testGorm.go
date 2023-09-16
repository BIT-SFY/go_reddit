package main

import (
	"reddit/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/reddit?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic("连接数据库失败!")
	}

	db.AutoMigrate(&models.User{})
}
