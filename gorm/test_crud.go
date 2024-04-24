package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
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
	db.AutoMigrate(&User{})
	//插入数据
	//user := User{Name: "dazhou", Age: 18, Birthday: time.Now()}
	users := []*User{
		{Name: "Jinzhu", Age: 18, Birthday: time.Now()},
		{Name: "Jackson", Age: 19, Birthday: time.Now()},
	}
	result := db.Create(&users)

	// 返回 error
	fmt.Printf("error:%v\n", result.Error)
	// 返回插入记录的条数
	fmt.Printf("rows:%d\n", result.RowsAffected)

}
