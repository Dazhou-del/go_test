package main

import "fmt"

func mayPanic() {
	err := new(error)
	if err != nil {
		fmt.Println("err !=nil", err)
	}
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main:", r)
		}
	}()

	mayPanic()
	fmt.Println("This line will not be executed.")
}
