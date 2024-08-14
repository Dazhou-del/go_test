package main

import (
	"errors"
	"fmt"
)

func main() {
	a()

}

func a() (resp string, err error) {
	if err = b(); err != nil {

		return
	}

	fmt.Println("a")
	return
}

func b() error {
	return errors.New("www")
}
