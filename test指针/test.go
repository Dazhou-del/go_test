package main

import (
	"fmt"
	"time"
)

func adds(a, b *int) {
	fmt.Println()
	fmt.Printf("a=%d,b=%d", *a, *b)
	*a = 67
	*b = 99
}

func addsTwo(a, b int) {
	fmt.Println()
	fmt.Printf("Two a=%d,b=%d", a, b)
	a = 67
	b = 99
}

func main() {
	var a *int
	c := 22
	d := 33
	a = &c                //将c的地址赋值给a
	fmt.Printf("a=%d", a) //a824633786504
	value := *a           //取这个地址的值
	fmt.Println()
	fmt.Printf("value=%d", value) //value=22
	addsTwo(c, d)
	adds(&c, &d)
	fmt.Println()

	fmt.Printf("c=%d,d=%d", c, d)
	//dateStr := time.Now().Format("2006-01-02")
	dateString := "2024-05-07"
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}

	// 计算并打印星期几
	weekday := date.Weekday()
	fmt.Println()
	fmt.Printf("%s is a %s\n", dateString, weekday.String())

	var acc uint64
	acc = 1
	fmt.Println()
	fmt.Printf("acc is a %d\n", acc)

	for i := 0; i < int(acc); i++ {
		fmt.Println("执行一次")
		//if acc == 1 {
		//	return
		//}
		//if i == int(acc)-1 {
		//	return
		//}
		if i == int(acc)-1 {
			fmt.Println("最后一次")

		}

	}
}
