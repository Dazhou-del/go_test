//package main
//
//import (
//	"fmt"
//	"strconv"
//)
//
//func uniqueInts(slice []int) []int {
//	// 创建一个 map 来存储唯一元素
//	seen := make(map[int]struct{})
//	var unique []int
//
//	for _, num := range slice {
//		// 使用 map 来检查元素是否已经存在
//		if _, exists := seen[num]; !exists {
//			// 如果元素是唯一的，则添加到 unique 切片中
//			unique = append(unique, num)
//			// 在 map 中标记这个元素已经存在
//			seen[num] = struct{}{}
//		}
//	}
//
//	return unique
//}
//
//func main() {
//	numbes := []int{}
//
//	// 去重后的数字列表
//	uniqueNumbers := uniqueInts(numbes)
//
//	// 输出去重后的数字列表
//	fmt.Println(len(uniqueNumbers))
//	fmt.Println(uniqueNumbers)
//
//	nuwws := []string{}
//	for _, value := range numbes {
//		nuwws = append(nuwws, strconv.Itoa(value)+",")
//	}
//	fmt.Println(nuwws)
//}
