package main

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/threading"
)

func main() {
	fmt.Println("start")
	chanStr := make(chan string, 1)
	chanStr <- "hello"
	threading.GoSafe(func() {
		val := <-chanStr
		fmt.Println("val", val)
	})
	fmt.Println("end")
}
