package main

import (
	"fmt"
	"github.com/samber/lo"
)

type ssc struct {
	age  int
	name string
	sex  bool
}

func main() {
	sscs := make([]ssc, 0)
	for i := 0; i < 5; i++ {
		newSsc := ssc{
			age:  i,
			name: "测试",
			sex:  true,
		}

		sscs = append(sscs, newSsc)
	}

	m := lo.SliceToMap(sscs, func(item ssc) (int, ssc) {
		return item.age, item
	})

	for i := 0; i < 5; i++ {
		if v, ok := m[i]; ok {
			fmt.Println(v.age, v.name)
		}
	}

	var cc []int
	ints := make([]int, 0)
	fmt.Println(ints)
	fmt.Println(cc)
	fmt.Println(cc == nil)
	fmt.Println(ints == nil)

}
