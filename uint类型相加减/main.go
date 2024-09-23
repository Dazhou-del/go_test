package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var c string

	c = "测试s"
	fmt.Println(len(c))
	fmt.Println(utf8.RuneCountInString(c))

	www := []uint64{1, 3, 4, 56, 7}
	for i, v := range www {
		fmt.Println("i:", i)
		fmt.Println("v:", v)
	}
}
