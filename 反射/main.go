package main

import (
	"fmt"
	"reflect"
)

type cat struct {
	Name string
}

func (*cat) asd() {
	fmt.Println("asd")
}

func (*cat) asds() {
	fmt.Println("asds")
}

func main() {
	//f := 2.3455
	//
	//typeOf := reflect.TypeOf(f)
	//valueOf := reflect.ValueOf(f)
	//fmt.Println(valueOf)
	//fmt.Println(valueOf.String())
	//fmt.Println(typeOf.String())
	//
	//fmt.Println("--------------------------------------------------")
	//
	//catw := cat{Name: "dazhou"}
	//catOf := reflect.TypeOf(catw)
	//catValueOf := reflect.ValueOf(catw)
	//fmt.Println(catValueOf)
	//fmt.Println(catValueOf.String())
	//fmt.Println(catOf.String())
	//userP := user{Name: "dazhou", age: 20, Married: true}
	//inspectStruct(userP)
	//inspectFunc(Add)

	//css := cat{Name: "dazhou"}
	//css.asd()
	//css.asds()
	//inspectStructs(cat{})
	//u := User{
	//	Name: "dj",
	//	Age:  18,
	//}
	//inspectMethod(&u)

	invoke(Add, 1, 2)
	invoke(Greeting, "dazhou")
}

type user struct {
	Name    string
	age     int
	Married bool
}

func inspectStruct(u any) {
	valueOf := reflect.ValueOf(u)
	fmt.Println(valueOf)

	field := valueOf.NumField()

	fmt.Println(field)
	fmt.Println(valueOf.Field(0))
	fmt.Println(valueOf.Kind())
	println(valueOf.Field(0).Kind())
	fmt.Println(valueOf.Field(0).Int())
}

func Add(a, b int) int {
	return a + b
}

func Greeting(name string) string {
	return "hello " + name
}

func inspectFunc(funcs interface{}) {
	valueOf := reflect.TypeOf(funcs)

	in := valueOf.NumIn() // 参数个数
	t := valueOf.In(0)
	t2 := valueOf.NumOut()
	t3 := valueOf.Out(0)

	fmt.Println(in)
	fmt.Println(t)
	fmt.Println(t2)
	fmt.Println(t3)
}

func inspectStructs(o any) {
	of := reflect.TypeOf(o)
	//method := of.NumMethod()
	field := of.NumField()
	fmt.Println(field)
	fmt.Println(of.Method(0))
}

func inspectMethod(o interface{}) {
	t := reflect.TypeOf(o)

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)

		fmt.Println(m)
	}
}

type User struct {
	Name string
	Age  int
}

func (u *User) SetName(n string) {
	u.Name = n
}

func (u *User) SetAge(a int) {
	u.Age = a
}

func invoke(f any, args ...interface{}) {
	v := reflect.ValueOf(f)

	argsV := make([]reflect.Value, 0, len(args))
	for _, arg := range args {
		// 获取参数
		argsV = append(argsV, reflect.ValueOf(arg))
	}

	rets := v.Call(argsV)
	fmt.Println("ret:")
	for _, ret := range rets {
		fmt.Println(ret.Interface())
	}
}
