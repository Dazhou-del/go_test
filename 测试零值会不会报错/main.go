package main

import (
	"fmt"
	"strconv"
)

type stu struct {
	age int
	ss  string
}

func main() {
	s := new(stu)
	s.ss = "0.00"
	s.age = 100

	price, err := strconv.ParseFloat(s.ss, 64)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(price)
}
