package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type testStu struct {
	sid  int64
	name int64
	age  int64
	nem  int64
	a    int64
	b    int64
	c    int64
	d    int64
	e    int64
	f    int64
	g    int64
	h    int64
	i    int64
	j    int64
	k    int64
	o    int64
	p    int64
	z    int64
	x    int64
	n    int64
}

type stu struct {
	sid  int
	name string
	age  int
	nem  int
}

type stu2 struct {
	id   int
	name string
	age  int
}

func main() {

	stus := make([]stu, 0, 5000)
	testStuList := make([]testStu, 0, 5000)
	for i := 0; i < 6000; i++ {
		var tem testStu
		tem.sid = int64(i)
		tem.name = int64(i)
		tem.age = int64(i)
		tem.nem = int64(i)
		tem.a = int64(i)
		tem.b = int64(i)
		tem.c = int64(i)
		tem.d = int64(i)
		tem.e = int64(i)
		tem.f = int64(i)
		tem.g = int64(i)
		tem.h = int64(i)
		tem.j = int64(i)
		tem.k = int64(i)
		tem.p = int64(i)
		tem.n = int64(i)
		testStuList = append(testStuList, tem)
	}

	// 单个元素的大小
	elemSize := unsafe.Sizeof(testStuList)

	// 切片背后数组的大小
	val := reflect.ValueOf(testStuList)
	sliceMemSize := val.Type().Size()

	cc := len(testStuList) * int(reflect.TypeOf(testStuList).Elem().Size())

	// 打印切片头的大小和数组的大小
	fmt.Printf("Header size: %d bytes\n", elemSize)
	fmt.Printf("Array size: %d bytes\n", sliceMemSize)
	fmt.Printf("cc size: %d bytes\n", cc)

	stus2 := make([]stu2, 0, 5000)

	var success []int

	for _, v := range stus {
		var ids []int
		for _, value := range stus2 {
			if v.sid == value.id {
				ids = append(ids, v.sid)
			}
		}

		if v.nem != len(ids) {
			success = append(success, v.sid)
		}
	}
}
