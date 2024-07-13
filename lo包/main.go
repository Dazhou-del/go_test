package main

import (
	"fmt"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"strconv"
	"strings"
	"time"
)

type student struct {
	name  string
	age   int
	class string
}

type class struct {
	name string
	stus []student
}

func main() {
	//loFilter()
	//loMap()
	//lopMap()
	//loFilterMap()
	//loFlatMap()
	//loReduce()
	//loReduceRight()
	//loFroEach()
	//lopForEach()
	//loTimes()
	//lopTimes()
	//lopGroupBy()
	//loChuck()
	loDrop()
}

func loFilter() {
	var a = []student{{age: 22, name: "dazhou", class: "3班"},
		{age: 24, name: "xiaom", class: "4班"},
		{age: 28, name: "xiaohua", class: "5班"},
		{age: 30, name: "xiaodong", class: "6班"}}

	filter := lo.Filter(a, func(i student, index int) bool {
		return i.age > 23
	})

	fmt.Println("a", a)
	fmt.Println("filter", filter)
}

func loMap() {
	b := []int64{1, 2, 3}

	fmt.Println("loMap开始时间:", time.Now())
	for i := 0; i < 1000000; i++ {
		_ = lo.Map(b, func(item int64, index int) string {
			return strconv.FormatInt(item, 10)
		})

		//fmt.Println("loMap.b", b)
		//fmt.Println("loMap.bStr", bStr)
	}

	fmt.Println("loMap结束:", time.Now())
}

func lopMap() {
	b := []int64{1, 2, 3}

	fmt.Println("lopMap开始时间:", time.Now())
	for i := 0; i < 1000000; i++ {
		_ = lop.Map(b, func(item int64, _ int) string {
			return strconv.FormatInt(item, 10)
		})

		//fmt.Println("lopMap.b", b)
		//fmt.Println("lopMap.bStr", bStr)
	}
	fmt.Println("lopMap结束时间:", time.Now())
}

func loFilterMap() {
	filterMap := lo.FilterMap([]string{"cpu", "gpu", "mouse", "keyboard"}, func(item string, _ int) (string, bool) {
		return strings.ToUpper(item), true
	})

	fmt.Println("filterMap:", filterMap)
}

func loFlatMap() {
	flatMap := lo.FlatMap([]int64{0, 1, 2}, func(item int64, index int) []string {
		return []string{
			strconv.FormatInt(item, 10),
			strconv.FormatInt(item, 10),
		}
	})

	fmt.Println("flatMap:", flatMap)
}

func loReduce() {
	sum := lo.Reduce([]int{1, 2, 3, 4}, func(agg int, item int, index int) int {
		return agg + item
	}, 0)

	fmt.Println("sum:", sum)
}

func loReduceRight() {
	right := lo.ReduceRight([][]int{{0, 1}, {2, 3}, {4, 5}}, func(agg []int, item []int, index int) []int {
		return append(agg, item...)
	}, []int{})

	fmt.Println("right:", right)

	s := lo.ReduceRight([]int{1, 2, 3, 4}, func(agg int, item int, index int) int {
		return agg + item
	}, 0)

	fmt.Println("s:", s)
}

func loFroEach() {
	lo.ForEach([]string{"hello", "world"}, func(x string, _ int) {
		println(x)
	})
}

func lopForEach() {
	lop.ForEach([]string{"hello", "world"}, func(x string, _ int) {
		println(x)
	})
}

func loTimes() {
	times := lo.Times(3, func(index int) int {
		return lo.ReduceRight([]int{1, 2, 3, 4}, func(agg int, item int, index int) int {
			return agg + item
		}, 0)
	})

	fmt.Println("times:", times)
}

func lopTimes() {
	times := lop.Times(3, func(index int) int {
		return lo.ReduceRight([]int{1, 2, 3, 4}, func(agg int, item int, index int) int {
			return agg + item
		}, 0)
	})

	fmt.Println("times:", times)
}

func lopGroupBy() {
	groupBy := lo.GroupBy([]int{1, 3, 5, 6}, func(index int) int {
		return index % 3
	})

	fmt.Println("groupBy:", groupBy)
}

func loChuck() {
	chunk := lo.Chunk([]int{0, 1, 2, 3, 4, 5}, 2)

	chunk2 := lo.Chunk([]int{0, 1, 2, 3, 4, 5, 6}, 2)

	chunk3 := lo.Chunk([]int{}, 2)

	chunk4 := lo.Chunk([]int{0}, 2)

	fmt.Println("chunk:", chunk)
	fmt.Println("chunk2:", chunk2)
	fmt.Println("chunk3:", chunk3)
	fmt.Println("chunk4:", chunk4)
}

func loDrop() {
	var a = []student{{age: 22, name: "dazhou", class: "3班"},
		{age: 24, name: "xiaom", class: "4班"},
		{age: 28, name: "xiaohua", class: "5班"},
		{age: 30, name: "xiaodong", class: "6班"}}

	drop := lo.Drop(a, 2)

	fmt.Println("a:", a)
	fmt.Println("drop:", drop)

}
