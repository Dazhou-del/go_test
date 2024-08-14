package main

import "fmt"

type stu struct {
	name string
	age  int
}

type stussss struct {
	stus []stu
	name string
	age  int
}

var stuList []stu

func main() {

	var stussssaaa stussss
	stus := a()
	stus2 := a()
	fmt.Println(stus)
	fmt.Println(stus2)

	stuscc := c()
	stuList := append(stuList, stuscc...)

	fmt.Println(stuList)

	var stuListaaa []stu
	stussssaaa.stus = stuListaaa

	fmt.Println(stussssaaa.stus)
}

func add(string2 string, age int) []stu {
	var stua stu
	stua.name = string2
	stua.age = age

	stuList = append(stuList, stua)

	return stuList
}

func a() []stu {
	s := "ss"
	age := 22

	stuList := add(s, age)

	return stuList
}

func c() []stu {
	var stus []stu
	var stua stu
	s := "ss"
	age := 22
	stua.name = s
	stua.age = age

	stus = append(stus, stua)

	return stus
}
