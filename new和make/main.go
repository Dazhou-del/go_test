package main

import "fmt"

func main() {
	i := new([]int)
	fmt.Println(i) // &[]
	//i2 := append(i, 12)
	//fmt.Println(i2)

	ints := make([]int, 0, 0) // []
	ints = append(ints, 12)
	fmt.Println(ints)

	int64s := make(chan int64)
	fmt.Println(int64s)
}
