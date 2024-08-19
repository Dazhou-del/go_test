package main

import "fmt"

func Counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	counter := Counter()
	fmt.Println(counter()) // 输出: 1
	fmt.Println(counter()) // 输出: 2
	fmt.Println(counter()) // 输出: 3
}
