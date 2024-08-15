package main

import "fmt"

type OrderState interface {
	next(order Order)
	previous(Order Order)
	printStatus()
}

type CreatedOrder struct {
	OrderState
}

type PaidState struct {
	OrderState
}

func (o *CreatedOrder) printStatus() {
	fmt.Println("Order created")
}

func (o *CreatedOrder) next(order Order) {
	fmt.Println("Order created")
}

func (o *CreatedOrder) previous(order Order) {
	fmt.Println("Order created")
}

func (o *PaidState) printStatus() {
	fmt.Println("Order PaidState")
}

func (o *PaidState) next(order Order) {
	fmt.Println("Order PaidState")
}

func (o *PaidState) previous(order Order) {
	fmt.Println("Order PaidState")
}

type Order struct {
}
