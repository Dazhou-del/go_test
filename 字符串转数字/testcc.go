package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "我的"
	a, err := strconv.Atoi(str)
	if err == nil {
		fmt.Println("字符串", str, "可以转换为整数")
	} else {
		fmt.Println("字符串", str, "不可以转换为整数")
	}

	fmt.Println("a", a)
}
