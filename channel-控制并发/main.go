package main

import "fmt"

func main() {
	channelList := make(chan int, 3)

	for {
		go func() {
			defer xx(channelList)
			channelList <- 1
			worker()

		}()
	}
}

func worker() {
	fmt.Println("正在执行任务")
}

func xx(channelList chan int) {
	<-channelList

}
