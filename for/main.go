package main

import "fmt"

func main() {

	number := 0

	for i := 0; i < 10; i++ {
		if i == 1 && number == 0 {
			fmt.Println(number)
			number++
		}

		if i%2 == 0 && number > 0 {
			fmt.Println(number)
			number++
		}
	}
}
