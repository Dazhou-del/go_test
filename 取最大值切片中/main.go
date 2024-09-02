package main

import (
	"fmt"
	"github.com/samber/lo"
)

// 定义Rental结构体
type Rental struct {
	ID           int
	total_rental float64
}

func main() {
	// 创建一些Rental实例
	rentals := []Rental{
		{ID: 1, total_rental: 1200.50},
		{ID: 2, total_rental: 1500.75},
		{ID: 3, total_rental: 1800.00},
	}

	// 使用samber/lo库找到total_rental最大的Rental
	maxRental := lo.MaxBy(rentals, func(r Rental, v Rental) bool {
		return r.total_rental > v.total_rental
	})

	fmt.Println(maxRental)
}
