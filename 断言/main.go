package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i any
	i = 22
	value, ok := i.(int)
	fmt.Println(value, ok)

	var x int
	x = 22
	fmt.Println(reflect.TypeOf(x), x)
}
