package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: received signal, exiting.\n", id)
			return
		default:
			fmt.Printf("Worker %d: running...\n", id)
			time.Sleep(500 * time.Millisecond) // 模拟工作
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// 启动10个goroutine
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(ctx, id)
		}(i)
	}

	// 等待一段时间，然后取消context
	time.Sleep(2 * time.Second)
	fmt.Println("Cancelling all workers...")
	cancel()

	// 等待所有goroutine结束
	wg.Wait()
	fmt.Println("All workers finished.")
}
