package main

import (
	"context"
	"fmt"
)

func main() {
	// 通过通道关闭
	// 如果 Goroutine 中使用了阻塞操作，如 I/O 操作或者等待通道，那么关闭通道可能无法停止 Goroutine。
	// 在这种情况下，可以使用 context.Context 来停止 Goroutine。
	/*	done := make(chan bool)

		go myGoroutine(done)

		time.Sleep(time.Second * 5)

		close(done)*/

	// 通过context关闭
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	//defer cancel()
	//go myGoroutine2(ctx)
	//time.Sleep(time.Second * 3)

}

// 通过通道关闭
func myGoroutine(done chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("hello world")
		}
	}
}

// 通过context取消
func myGoroutine2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("hello world22")
		}
	}

}
