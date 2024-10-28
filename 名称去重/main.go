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

	valueMap["测试4"] = map[int]string{1: "租满6个月后可随租随还"}
	fmt.Println("valueMap", valueMap)

	for _, innerMap := range valueMap {
		for i2 := range innerMap {
			if i2 == 1 {
				innerMap[i2] = "测试过程"
			}
		}
	}

	fmt.Println("修改后的", valueMap)

	valueMap["测试4"] = map[int]string{1: "租满6个月后可随租随还"}

	var number int
	number++
	if valueMap["测试4"][1] != "" {
		valueMap["测试4"] = map[int]string{1: "租满6个月后可随租随还" + strconv.Itoa(number-1)}
	}

	fmt.Println("valueMap2", valueMap)

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
