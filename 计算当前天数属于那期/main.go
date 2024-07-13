package main

import (
	"fmt"
	"strconv"
)

func main() {

	//var startDay, endDay time.Time
	////2023-10-21 00:00:00
	//timeStart := int64(1697817600)
	////2024-07-02
	//timeEnd := int64(1719925310)
	////now := time.Now().Unix()
	//now := int64(1721577600)
	//keys := make([]int, 100)
	//
	//var j int
	//for i := 1; i < len(keys); i++ {
	//	if i == 1 {
	//		startDay = time.Unix(timeStart, 0)
	//	} else {
	//		startDay = endDay.AddDate(0, 0, +1)
	//	}
	//
	//	if time.Unix(timeStart, 0).AddDate(0, 1, 0).Unix() > time.Unix(timeEnd, 0).Unix() {
	//		endDay = time.Unix(timeEnd, 0)
	//	} else {
	//		endDay = time.Unix(timeStart, 0).AddDate(0, i, -1)
	//	}
	//
	//	if now < startDay.Unix() {
	//		j++
	//
	//		break
	//	} else if now >= startDay.Unix() && now <= endDay.Unix() {
	//		fmt.Println(i)
	//
	//		break
	//	}
	//	fmt.Println("startDay", startDay)
	//	fmt.Println("endDay", endDay)
	//
	//	fmt.Println("j", j)
	//	fmt.Println("i", i)
	//
	//}
	//
	//// 大于12期的直接取最后一期
	//if now > timeEnd {
	//	fmt.Println("大于12期")
	//}

	a := 12.31
	//// 'g' 格式会根据数值选择最紧凑的表示形式，可能是固定小数点或科学记数法
	//s := strconv.FormatFloat(f, 'g', -1, 64)
	//fmt.Println(s) // 输出：123.456

	// 'f' 格式总是使用固定小数点表示形式
	// 第三个参数是小数点后要保留的位数
	s := strconv.FormatFloat(a, 'f', 2, 64)
	fmt.Println(s) // 输出：123.46

}
