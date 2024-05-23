package main

import (
	"fmt"
	"time"
)

func main() {
	// 示例时间戳，单位为秒
	//当前时间
	startTime := int64(1715849973)

	//认证时间
	renzhengTime := int64(1707898363)

	if renzhengTime < startTime {
		Time := time.Unix(startTime, 0)
		fmt.Printf("当前时间：", Time)

		rengzTime := time.Unix(renzhengTime, 0)
		fmt.Printf("认证时间：", rengzTime)

		diff60 := Time.Sub(rengzTime)
		days60 := diff60.Hours() / 24
		fmt.Println("两个时间相差多少天,", days60)

		if days60 < 60 {
			fmt.Println("在60天内")
		} else {
			fmt.Println("不在在60天内")
		}

		if days60 < 90 {
			fmt.Println("在90天内")
		} else {
			fmt.Println("不在90天内")
		}
	} else {
		fmt.Println("新在60天内")
	}

}
