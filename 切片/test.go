package main

import "fmt"

func Test_slice() {

}

func main() {
	ints := make([]int, 10)
	ints = append(ints, 2)

	fmt.Printf("ints:%v len is:%d cap %d", ints, len(ints), cap(ints)) //ints:[0 0 0 0 0 0 0 0 0 0 2] len is:11 cap 20
	fmt.Println()

	i := make([]int, 0, 11)
	i = append(i, 3)
	fmt.Printf("i:%v len is:%d cap %d", i, len(i), cap(i)) //i:[3] len is:1 cap 11
	fmt.Println()

	i2 := make([]int, 10, 12)
	i3 := i2[8:]
	fmt.Printf("i3:%v len is:%d cap: %d", i3, len(i3), cap(i3)) //i3:[0 0] len is:2 cap 4
	fmt.Println()

	s := make([]int, 10, 12)
	s1 := s[8:]
	s1[0] = -1
	fmt.Printf("s: %v", s)
	fmt.Println()

	fmt.Println()
	i4 := make([]int, 10, 12)
	i5 := i4[8:]
	i5[0] = -1
	fmt.Printf("i5:%v len is:%d cap: %d", i5, len(i5), cap(i5)) //i5:[-1 0] len is:2 cap: 4
	fmt.Println()

	/*	i6 := make([]int, 10, 12)
		i7 := i6[10] //make声明并初始化，长度为10，容量为12，访问下标为10，大于长度访问越界
		//是否越界
		fmt.Println(i7)*/

	//添加切片元素
	i6 := make([]int, 0, 10)
	i7 := append(i6, []int{1, 3, 5}...)
	fmt.Printf("i7:%v len is:%d cap: %d", i7, len(i7), cap(i7)) //i7:[0 0 0 0 0 1 3] len is:7 cap: 10
	fmt.Println()

	//删除切片元素
	i7 = append(i7[:1], i7[2:]...)                              //总结一下就是：要从切片a中删除索引为index的元素，操作方法是a = append(a[:index], a[index+1:]...)
	fmt.Printf("i7:%v len is:%d cap: %d", i7, len(i7), cap(i7)) //i7:[0 0 0 0 0 1 3] len is:7 cap: 10
	fmt.Println()

	i8 := []int{1, 2, 3, 4, 5}
	i9 := i8[1:]
	fmt.Printf("i9:%v len is:%d cap: %d", i9, len(i9), cap(i9))
	fmt.Println()

	//切片变量
	for index, value := range i8 {
		fmt.Printf("index:%d,value%d", index, value)
		fmt.Println()
	}
	fmt.Println("错误代码")
}
