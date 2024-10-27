package main

import (
	"fmt"
	"time"
)

func main() {
	ints := make(chan int, 100)
	go worker(ints)

	// 阻塞任务
	for i := 0; i < 10; i++ {
		ints <- i
	}

	select {
	case <-time.After(time.Second * 20):
	}
}

func worker(taskCh <-chan int) {
	const N = 5

	//启动协程工作
	for i := 0; i < 5; i++ {
		go func(id int) {
			for {
				task := <-taskCh
				fmt.Printf("finish task: %d by worker %d\n", task, id)
				time.Sleep(time.Second)
			}
		}(i)
	}
}
