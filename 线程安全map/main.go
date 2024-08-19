package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.Map

	// 写入
	m.Store("key1", 22)
	m.Store("key2", 33)

	// 读取
	age, _ := m.Load("key1")
	fmt.Println(age.(int))

	// 遍历
	m.Range(func(key, value interface{}) bool {
		name := key.(string)
		gee := value.(int)
		fmt.Println(name, gee)
		return true
	})

	// 删除
	m.Delete("key1")
	age, ok := m.Load("key1")
	fmt.Println(age, ok)

	// 读取或写入
	m.LoadOrStore("stefno", 100)
	age, _ = m.Load("stefno")
	fmt.Println(age)

	ch := make(chan string, 6)
	go func() {
		for {
			v, ok := <-ch
			if !ok {
				fmt.Println("结束")
				return
			}
			fmt.Println(v)
		}
	}()

	ch <- "煎鱼还没进锅里..."
	ch <- "煎鱼进脑子里了！"
	close(ch)
	time.Sleep(time.Second)

}
