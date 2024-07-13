package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

func main() {
	// 假设这是从数据库获取的JSON字符串
	jsonDataFromDB := `["75421.21","94375.2123","83133","93722"]`

	// 创建一个字符串切片来存放解码后的数据
	var strNumbers []string

	// 解码JSON字符串到strNumbers
	if err := json.Unmarshal([]byte(jsonDataFromDB), &strNumbers); err != nil {
		log.Fatal("Error during Unmarshal:", err)
	}

	// 如果需要将这些字符串转换为整数切片
	var intNumbers []float64
	for _, strNum := range strNumbers {
		num, err := strconv.ParseFloat(strNum, 64)
		if err != nil {
			log.Printf("Error converting '%s' to int: %v", strNum, err)
			continue // 如果转换失败，可以选择跳过或处理错误
		}
		intNumbers = append(intNumbers, num)
	}

	fmt.Println("String slice:", strNumbers)
	fmt.Println("Integer slice:", intNumbers)

}
