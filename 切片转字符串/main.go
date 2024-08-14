package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 定义一个整数数组
	arr := [5]int{1, 2, 3, 4, 5}

	// 将数组转换为字符串
	strArr := arrayToString(arr[:])

	// 打印结果
	fmt.Println(strArr)
}

// arrayToString 是一个辅助函数，用于将整数切片转换为字符串
func arrayToString(slice []int) string {
	var str string
	for i, v := range slice {
		if i == 0 {
			str += strconv.Itoa(v)
		} else {
			str += "," + strconv.Itoa(v)
		}
	}
	return str
}
