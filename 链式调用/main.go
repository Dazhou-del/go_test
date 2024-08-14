package main

import "fmt"

type stu struct {
	age  int
	name string
}

func (s *stu) AddAge() *stu {
	s.age++

	return s
}

func (s *stu) setName(name string) {
	s.name = name
}

func main() {
	s := stu{name: "da", age: 20}
	fmt.Println("开始s", s)

	s.AddAge().setName("ssw")
	fmt.Println(s)
}
