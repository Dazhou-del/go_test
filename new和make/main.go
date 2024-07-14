package main

import "fmt"

func main() {
	i := new([]int)
	fmt.Println(i) // &[]

	ints := make([]int, 0, 0) // []
	fmt.Println(ints)

	int64s := make(chan int64)
	fmt.Println(int64s)
}
