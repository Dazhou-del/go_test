package main

import (
	"fmt"
	"time"
)

func main() {
	bools1 := make(chan bool)
	bools2 := make(chan bool)
	bools3 := make(chan bool)
	// 声明为nil的通道
	var xsx chan bool
	fmt.Println(xsx)

	go func() {
		<-bools1 // 起点 外部设置一值时才能调用
		fmt.Println("goroutine1")
		// 调用第二个goroutine
		bools2 <- true
	}()

	go func() {
		<-bools2
		fmt.Println("goroutine2")
		// 调用第一个goroutine
		bools3 <- true
	}()

	go func() {
		<-bools3
		fmt.Println("goroutine3")
		// 调用第一个goroutine
		bools1 <- true
	}()

	bools1 <- true // 起点调用

	// 需要等待两个goroutine调用 如果主goroutine结束两个并不会被调用
	time.Sleep(time.Second * 5)
}
