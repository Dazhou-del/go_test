package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().AddDate(0, 1, 17)

	// 获取当前月份的第一天
	currentMonthFirstDay := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

	// 获取上个月的第一天
	lastMonthFirstDay := currentMonthFirstDay.AddDate(0, -1, 0)
	// 上个月的第10天
	lastMonthTenth := lastMonthFirstDay.AddDate(0, 0, 10-1)
	// 当前月的第10天
	thisMonthTenth := currentMonthFirstDay.AddDate(0, 0, 10-1)

	// 检查当前时间是否在上个月的第10天到这个月的第10天之间
	if lastMonthTenth.Before(now) && thisMonthTenth.After(now) {
		time := time.Now()
		dateStr := time.Format("1")
		fmt.Println(dateStr)

		fmt.Println("当前时间在上个月10号到这个月10号之间month", dateStr)
	} else {
		// 如果不在这个区间内，那么当前时间属于下个月
		nextMonthFirstDay := currentMonthFirstDay.AddDate(0, 1, 0)
		fmt.Printf("当前时间不属于上述区间，属于 %s 月份\n", nextMonthFirstDay.Format("2006-01"))
	}
}
