package main

import "fmt"

func main() {
	output := make(chan int)
	defer func() {
		for range output {
			panic("more than one element written in reducer")
		}
	}()
	output <- 1

	// 从通道接收值
	go func() {
		for i := range output {
			fmt.Println("Received value:", i)
		}
	}()
	close(output)
}
