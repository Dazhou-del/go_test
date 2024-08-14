package main

import (
	"fmt"
	"time"
)

func main() {
	dree := make(map[string]map[int64]int64)
	startTime := time.Now().Unix()
	//如果认证时间大于前一天时间 那么肯定是新商家
	VerifyTime := int64(1715083235)
	if VerifyTime >= startTime {
		// 在外层 map 中创建一个新的内层 map
		dree["new"] = make(map[int64]int64)

		dree["new"][22222] = 1

		dree["new_three"] = make(map[int64]int64)

		dree["new_three"][22222] = 1
		fmt.Println("大于新商家")
	}

	//认证时间小于前一天时间 计算在多少天内
	startTimes := time.Unix(startTime, 0)
	verifyTimes := time.Unix(VerifyTime, 0)
	diff := startTimes.Sub(verifyTimes)
	days := diff.Hours() / 24
	dree["new_three"] = make(map[int64]int64)
	dree["new"] = make(map[int64]int64)
	//商家入驻审核通过后2个月内
	if days < 60 {

		dree["new"][22222] = 1

		dree["new_three"][22222] = 1
		fmt.Println("新商家")
	}

	//商家入驻审核通过后3个月内
	if days < 90 {

		dree["new_three"][22222] = 1
	}

	if _, ok := dree["new"][22222]; !ok {
		// 在外层 map 中创建一个新的内层 map
		dree["old"] = make(map[int64]int64)

		dree["old"][22222] = 1
	}

	fmt.Println(dree)
}
