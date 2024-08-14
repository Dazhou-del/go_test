package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 向通道发送数据
func sendData(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i
		time.Sleep(time.Second) // 模拟耗时操作
	}
	close(ch)
}

// 从通道接收数据
func receiveData(ch <-chan int) {
	wg := sync.WaitGroup{}
	wg.Add(5)
	for num := range ch {
		go func() {
			defer wg.Done()
			fmt.Println("Received:", num)
		}()
	}
	//wg.Wait()
}

func main() {
	intChannels := [3]chan int{make(chan int, 1), make(chan int, 1), make(chan int, 1)}

	index := rand.Intn(3)
	fmt.Println(index)
	intChannels[index] <- index
	intChannels[index] <- index

	select {
	case <-intChannels[0]:
		fmt.Println("The first candidate case is selected.")
	case <-intChannels[1]:
		fmt.Println("The second candidate case is selected.")
	case elem := <-intChannels[2]:
		fmt.Printf("The third candidate case is selected, the element is %d.\n", elem)
	default:

	}

	fmt.Println("No candidate case is selected!")

	//dataChan := make(chan int) // 创建一个普通的双向通道
	//
	//// 启动goroutine发送数据到通道
	//go sendData(dataChan)
	//
	//// 使用select语句从通道中接收数据
	//select {
	//case <-time.After(3 * time.Second):
	//	fmt.Println("Timeout occurred. Exiting...")
	//default:
	//	receiveData(dataChan) // 直接调用receiveData函数，从通道中接收数据
	//}

}
