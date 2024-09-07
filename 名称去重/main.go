package main

import (
	"fmt"
	"github.com/samber/lo"
	"strconv"
)

func main() {
	name := []string{"测试1", "测试1", "测试1", "测试1", "测试1", "测试2", "测试3", "测试4"}
	uniq := lo.Uniq(name)
	var i = 0
	valueMap := make(map[string]map[int]string, 40)

	valueMap = lo.SliceToMap(uniq, func(item string) (string, map[int]string) {
		i++
		if i == 1 {

			return item, map[int]string{0: "租满6个月后可随租随还"}
		}

		return item, map[int]string{0: "租满6个月后可随租随还" + strconv.Itoa(i-1)}
	})

	if aw, ok := valueMap["测试2"]; ok {
		if valueMap["测试2"][0] != "" {
			valueMap["测试2"] = map[int]string{1: valueMap["测试2"][0]}
		}
		fmt.Println(aw)
	}
	fmt.Println(valueMap["测试2"])

	fmt.Println("------------------")
	if aw, ok := valueMap["测试2"]; ok {
		if valueMap["测试2"][0] != "" {
			valueMap["测试2"] = map[int]string{1: valueMap["测试2"][0]}
		}

		fmt.Println(aw)
	}
	fmt.Println(valueMap["测试2"])

	fmt.Println("------------------")
	if aw, ok := valueMap["测试2"]; ok {

		if valueMap["测试2"][0] != "" {
			valueMap["测试2"] = map[int]string{1: valueMap["测试2"][0]}
		}
		fmt.Println(aw)
	}
	fmt.Println(valueMap["测试2"])

	if aw, ok := valueMap["测试1"]; ok {
		var mapc map[int]string
		mapc = map[int]string{1: valueMap["测试1"][0]}
		valueMap["测试1"] = mapc
		fmt.Println(aw)
	}
	fmt.Println(valueMap["测试1"])
}
