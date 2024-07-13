package main

import (
	"fmt"
	"strconv"
)

func main() {

	numbers := "5555,33333"
	atoi, err := strconv.Atoi(numbers)
	if err != nil {
		fmt.Printf("error:%v", err)
		fmt.Println()
	}

	fmt.Printf("atoi:%d", atoi)
}
