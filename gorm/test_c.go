package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Users struct {
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time
}

func main() {
	dsn := "root:dazhou520@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	//迁移 保证数据库表结构与结构体一样
	db.AutoMigrate(&Users{})
	var u Users
	db.First(u, 1)
	fmt.Printf("result: %v\n", u)
}
