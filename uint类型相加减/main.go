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

}
