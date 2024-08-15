package main

import "fmt"

func mayPanic() {
	err := new(error)
	if err != nil {
		fmt.Println("err !=nil", err)
	}
}

func main() {
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println("Recovered in main:", r)
	//	}
	//}()
	//
	//mayPanic()
	//fmt.Println("This line will not be executed.")
	i := 6

	name := funcName(i)
	fmt.Println("name:", name)
}

func funcName(int2 int) bool {
	if int2 != 4 && int2 != 5 {
		return false
	}

	return true
}
