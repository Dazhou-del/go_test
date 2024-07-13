package main

import (
	"fmt"
	"strconv"
)

func main() {

	var a float64
	a = 12.311
	//// 'g' 格式会根据数值选择最紧凑的表示形式，可能是固定小数点或科学记数法
	//s := strconv.FormatFloat(f, 'g', -1, 64)
	//fmt.Println(s) // 输出：123.456

	// 'f' 格式总是使用固定小数点表示形式
	// 第三个参数是小数点后要保留的位数
	s := strconv.FormatFloat(a, 'f', 3, 64)
	fmt.Println(s) // 输出：123.46
}
