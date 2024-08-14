package main

import (
	"context"
	"fmt"
)

type animal struct {
	ctl context.Context
}

type function interface {
	add()
}

type cat struct {
	animal
	name string
}

type dog struct {
	animal
	name string
}

func (l *cat) add() {
	fmt.Println(l.name)
}

func (l *dog) add() {
	fmt.Println(l.name)
}

func main() {
	types := "cat"
	var a animal

	a.funcName(types).add()

}

func (l *animal) funcName(types string) function {
	switch types {
	case "cat":
		return &cat{name: "dazhou"}
	}
	return nil
}
