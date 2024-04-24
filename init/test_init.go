package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 定义一个全局变量db
// var db *sql.DB
//
// // 定义一个初始化数据库的函数
//
//	func initDb() (err error) {
//		//DSN:Data Source data
//		dsn := "root:dazhou520@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True"
//		//不会校验账号密码是否正确
//		// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
//		db, err = sql.Open("mysql", dsn)
//		if err != nil {
//			return err
//		}
//		err = db.Ping()
//		if err != nil {
//			return err
//		}
//		return nil
//	}
func main() {
	err := initDb()
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
}
