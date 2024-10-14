package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	// 获取当前时间并格式化
	formattedTime := time.Now().Format("2006010215040501")

	// 提取前4位和后6位
	prefix := formattedTime[:4]
	suffix := formattedTime[10:]

	// 组合前4位和后6位
	result := prefix + suffix

	// 打印结果
	fmt.Println(result)
}

func RandomNumber(n int) string {
	tmp := "%0" + strconv.FormatInt(int64(n), 10) + "v"

	return fmt.Sprintf(tmp, rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}
