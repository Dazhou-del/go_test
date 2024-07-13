package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "1"
	strSlice := strings.Split(str, ",")
	var intSlicse []int64
	strconvINt(strSlice, intSlicse)

	fmt.Println(strSlice)

}

func strconvINt(strSlice []string, intSlice []int64) {
	for _, s := range strSlice {
		num, err := strconv.Atoi(s)

		if err != nil {
			continue // 如果转换出错，可以选择处理错误或跳过
		}

		intSlice = append(intSlice, int64(num))
	}
}
