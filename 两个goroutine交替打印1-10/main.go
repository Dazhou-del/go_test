package main

import (
	"fmt"
	"sync"
)

func g1(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		fmt.Println("g1", i)
		ch <- i

		<-ch
	}
}

func g2(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		<-ch
		fmt.Println("g2:", i)
		ch <- i
	}
}

func main() {
	ints := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go g1(ints, &wg)
	go g2(ints, &wg)

	wg.Wait()

}
