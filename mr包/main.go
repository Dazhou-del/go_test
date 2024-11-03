package main

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/mr"
	"testing"
	"time"
)

func main() {
	//TestFinish()

	//TestForEach()

	TestMapReduce()
}

func TestMapReduce() {
	defer TrackTime(time.Now())
	data := []int{1, 2, 3, 4, 5}

	result, err := mr.MapReduce(func(source chan<- int) {
		for _, i := range data {
			source <- i
		}
	}, func(item int, writer mr.Writer[int], cancel func(error)) {
		time.Sleep(time.Second * 1)
		writer.Write(item * item)
	}, func(pipe <-chan int, writer mr.Writer[[]int], cancel func(error)) {
		var result []int
		for i := range pipe {
			result = append(result, i)
		}

		writer.Write(result)
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestForEach() {
	defer TrackTime(time.Now())
	data := []int{1, 2, 3, 4, 5}

	mr.ForEach(func(source chan<- int) {
		for _, i := range data {
			source <- i
		}
	}, func(item int) {
		time.Sleep(time.Second * 1)
		str := fmt.Sprintf("%d * %d = %d", item, item, item*item)
		fmt.Println(str)
	})
}

func TestFinish() {
	defer TrackTime(time.Now())

	err := mr.Finish(func() error {
		println("hello")
		time.Sleep(time.Second * 1)

		return nil
	}, func() error {
		println("hello2")
		time.Sleep(time.Second * 1)

		return nil
	}, func() error {
		println("hello3")
		time.Sleep(time.Second * 1)

		return nil
	})

	if err != nil {
		panic(err)
	}
}

func TrackTime(pre time.Time) time.Duration {
	elapsed := time.Since(pre)
	fmt.Println("elapsed:", elapsed)

	return elapsed
}

func TestTrackTime(t *testing.T) {
	defer TrackTime(time.Now()) // <-- 就是这里

	time.Sleep(500 * time.Millisecond)
}
