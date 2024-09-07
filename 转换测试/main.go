package main

import (
	"fmt"
	"strconv"
)

func main() {
	totalRental := "0.00"
	maxTotalRental, err := strconv.ParseFloat(totalRental, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(maxTotalRental)

	if maxTotalRental == 0 {
		fmt.Println("No rental")
	}
}
