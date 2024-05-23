package main

import (
	"fmt"
	"strings"
)

func main() {

	numbers := ``

	// 使用strings.Split将字符串按换行符分割成切片
	lines := strings.Split(numbers, "\n")

	// 使用strings.Builder来构建结果字符串
	var sb strings.Builder

	// 遍历切片，并在每个数字后面添加逗号，除了最后一个
	for i, line := range lines {
		sb.WriteString(line)
		if i < len(lines)-1 {
			sb.WriteString(",")
		}
	}

	// 获取构建的结果字符串
	result := sb.String()

	// 打印结果
	fmt.Println(result)

}
