package main

import "fmt"

func main() {

	var aa []int
	aa = append(aa, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21)
	ints := aa[0:5]
	fmt.Println(ints)
	fmt.Println(ints[5-1])
	cc := aa[5:10]
	fmt.Println(cc)

	var cad float64
	cad = float64(5580.0) / 1000.0
	fmt.Println(cad)
}
