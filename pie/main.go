package main

import (
	"fmt"
	"github.com/elliotchance/pie/v2"
)

func main() {
	a := []int{1, 2, 3}
	b := []int{1, 4, 5, 6, 7}

	added, removed := pie.Diff(a, b)

	fmt.Println("added", added)
	fmt.Println("removed", removed)

	b2 := pie.Any(a, func(value int) bool {
		if value == 5 {
			return true
		}
		return false
	})
	fmt.Println("b", b2)
}
