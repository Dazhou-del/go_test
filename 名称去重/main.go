package main

import (
	"fmt"
	"github.com/samber/lo"
	"strconv"
)

func main() {
	name := []string{"测试1", "测试1", "测试1", "测试1", "测试1", "测试2", "测试3", "测试4"}
	fmt.Println(name)
	uniq := lo.Uniq(name)
	fmt.Println(uniq)
	var i = 0

	sliceToMap := lo.SliceToMap(uniq, func(item string) (string, string) {
		i++
		if i == 1 {
			return item, "租满6个月后可随租随还"
		}

		return item, "租满6个月后可随租随还" + strconv.Itoa(i-1)
	})
	fmt.Println(sliceToMap)

	if aw, ok := sliceToMap["测试2"]; ok {
		fmt.Println(aw)
	}

}
