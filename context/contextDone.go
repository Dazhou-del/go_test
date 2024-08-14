package main

import (
	"context"
	"fmt"
	"time"
)

// 假设这是你的Value类型
type Value int

func ContextDone(ctx context.Context, out chan<- Value) error {

	for {
		v, err := AllenHandler(ctx)

		if err != nil {
			return err
		}

		select {
		case <-ctx.Done():
			fmt.Printf("context has done")
			return ctx.Err()
		case out <- v:
		}

	}
}

func AllenHandler(ctx context.Context) (Value, error) {
	// ... 实现细节 ...
	return Value(42), nil
}

func main() {
	// 创建带有超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	// 创建输出通道
	out := make(chan Value)

	// 启动一个新的goroutine来调用ContextDone
	go func() {
		defer close(out) // 当goroutine结束时关闭通道
		err := ContextDone(ctx, out)
		if err != nil {
			fmt.Printf("ContextDone exited with error: %v", err)
		}
	}()

	// 监听输出通道
	for v := range out {
		fmt.Printf("Received value: %d", v)
	}

	// 取消上下文，以停止ContextDone函数
	cancel()
}
