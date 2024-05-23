package main

import "fmt"

func main() {
	integers := []int{}
	// 创建一个新的字符串切片来保存带双引号的字符串
	quotedStrings := make([]string, len(integers))

	// 循环整数切片，为每个整数创建带双引号的字符串
	for i, num := range integers {
		quotedStrings[i] = "'" + fmt.Sprint(num) + "'" + ","
	}

	// 打印带双引号的字符串切片
	fmt.Println(quotedStrings)
}
