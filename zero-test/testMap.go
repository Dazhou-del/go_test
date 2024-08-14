package main

import (
	"fmt"
	"time"
)

func main() {
	//generateFunc := func(source chan<- int) {
	//	for i := 0; i < 10; i++ {
	//		source <- i
	//	}
	//}
	//
	//mapperFunc := func(item int) {
	//	fmt.Println("processing item", item)
	//}
	//
	//mr.ForEach(generateFunc, mapperFunc, mr.WithWorkers(4))
	// 假设这是两个Unix时间戳（秒）
	timestamp1 := int64(1721120771)
	timestamp2 := int64(1721811971) // 这个时间戳与第一个相差刚好不到7天

	// 将时间戳转换为time.Time类型
	time1 := time.Unix(timestamp1, 0)
	time2 := time.Unix(timestamp2, 0)

	// 计算两个时间之间的差值
	diff := time1.Sub(time2).Abs()

	// 检查差值是否小于7天
	if diff.Hours()/24 < 7 {
		fmt.Println("两个时间戳相差小于7天")
	} else {
		fmt.Println("两个时间戳相差大于或等于7天")
	}
}
