package main

import (
	"fmt"
	"time"
)

func main() {
	start := int64(1720108800)
	end := int64(1720454399)

	s := end - start
	fmt.Println(s)
	bill := float64((end - start) / 86400)
	fmt.Println("bill:", bill)

	startTime := time.Unix(start, 0)
	endTime := time.Unix(end, 0)

	deffTime := endTime.Sub(startTime).Abs()
	aa := deffTime.Hours() / 24

	fmt.Println(aa)
	// 订单投保生效时间小于7天
	if deffTime.Hours()/24 < 7 {
		fmt.Println("小于7")
	} else {
		fmt.Println("大于7")
	}
}
