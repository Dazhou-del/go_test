package main

import (
	"fmt"
	"time"
)

type singn struct {
	now  func() time.Time
	name string
}

func newSingn() *singn {
	return &singn{
		now:  time.Now,
		name: "dazhou",
	}
}

var _time time.Time

func (s *singn) fmtName() string {
	return s.name
}

func main() {
	//
	//_time = time.Now()
	//fmt.Println(_time)
	//_time = time.Now().AddDate(0, 0, 1)
	//fmt.Println(_time)
	//
	//s := newSingn()
	//fmt.Println("s.now", s.now())
	//fmt.Println("s.name:", s.fmtName())

	var aa []int
	cc := []int{}
	if aa == nil {
		fmt.Println("aa==nil")
	}

	if cc == nil {
		fmt.Println("cc==nil")
	}

	if len(cc) == 0 {
		fmt.Println("cc==0")
	}
}
