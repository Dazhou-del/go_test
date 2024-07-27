package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/threading"
	"sync"
	"time"
)

var i int
var wg sync.WaitGroup

func add() {
	//var m sync.Mutex // 互斥锁
	ints := make([]int, 0, 11111)

	for i := 0; i < 50; i++ {
		//m.Lock()
		ints = append(ints, i)
		//m.Unlock()
	}
	time.Sleep(1 * time.Second)
	fmt.Println(ints)

	fmt.Println("-----------------------------------------------------")

	i++

	defer wg.Done()
}

var db *sql.DB

func initDb() (err error) {
	//DSN:Data Source data
	dsn := "root:dazhou520@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True"
	//不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func insertData(is int) {
	ints := make([]int, 0, 150)
	//for i := 6000; i < is; i++ {
	ints = append(ints, i)
	sqlStr := "insert into book(title,author,price) values (?,?,?)"
	ret, err := db.Exec(sqlStr, "王五", "王五", is)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
	}
	fmt.Println(ret)
	//}
	fmt.Println("ints", ints)

}

type Book struct {
	id     int
	title  string
	author string
	price  int
}

func queryInsert() {
	ints := make([]int, 0, 150)
	//for i := 6000; i < is; i++ {
	querySql := "select * from book where  id=? "
	var book Book
	err := db.QueryRow(querySql, 222).Scan(&book.id, &book.title, &book.author, &book.price)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
	}

	ints = append(ints, i)
	sqlStr := "insert into book(title,author,price) values (?,?,?)"
	ret, err := db.Exec(sqlStr, book.author, book.title, 9999)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
	}
	fmt.Println(ret)
	//}
	fmt.Println("ints", ints)

}

func main() {
	/*	now := time.Now()
		fmt.Println("now", now)

		fmt.Println("aacc")
		wg.Add(50)
		for i := 0; i < 50; i++ {
			threading.GoSafe(func() {
				//defer wg.Done()
				add()
			})
			//add()
		}
		//time.Sleep(5 * time.Second)
		wg.Wait()

		fmt.Println("ddddd")

		fmt.Println("i:", i)

		end := time.Now()
		fmt.Println("end", end)*/

	now := time.Now()
	fmt.Println("now", now)
	err := initDb()
	if err != nil {
		fmt.Println(err)
	}

	//insertData(7000)
	/*	wg.Add(1)
		threading.GoSafe(func() {

			for i := 5000; i < 6000; i++ {
				insertData(i)
			}
			defer wg.Done()
		})
		wg.Wait()*/
	//wg.Add(100)
	//for i := 9600; i < 9700; i++ {
	//	//var m sync.Mutex // 互斥锁
	//	//m.Lock()
	//	threading.GoSafe(func() {
	//		defer wg.Done()
	//		insertData(i)
	//	})
	//	//m.Unlock()
	//}
	//wg.Wait()
	//end := time.Now()
	//fmt.Println("end", end)
	//wg.Add(10)
	//for i := 0; i < 10; i++ {
	//wg.Add(1)
	threading.GoSafe(func() {
		fmt.Println("我先执行")
		//defer wg.Done()
		insertData(3333)
	})
	//}

	//threading.GoSafe(func() {
	//	queryInsert()
	//})

	//
	//threading.GoSafe(func() {
	//	time.Sleep(3 * time.Second)
	//	fmt.Println("我后执行")
	//
	//	defer wg.Done()
	//	queryInsert()
	//})

	//wg.Wait()

	end := time.Now()
	fmt.Println("end", end)
	fmt.Println("结束")

}
