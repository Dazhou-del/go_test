package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	//numbe2 := []int{1, 2, 3}

	//maps := map[int]string{1: "a", 2: "b", 3: "c"}

	contain, err := Contain(2, numbers)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(contain)
}

func Contain(obj interface{}, target interface{}) (bool, error) {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true, nil
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		}
	}

	return false, errors.New("not in array")
}
