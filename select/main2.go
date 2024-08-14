package main

import (
	"fmt"
	"time"
)

// 商品结构体
type Product struct {
	ID       int
	Name     string
	Quantity int
}

// 从供应商获取商品信息
func fetchProductInfo(sendCh chan<- Product) {
	// 模拟从供应商获取商品信息的耗时操作
	time.Sleep(2 * time.Second)

	// 模拟获取到的商品信息
	product := Product{
		ID:       1,
		Name:     "Product A",
		Quantity: 100,
	}

	// 发送商品信息到通道
	sendCh <- product
}

// 处理订单并更新库存
func processOrder(recvCh <-chan Product, orderID int) {
	// 从通道接收商品信息
	product := <-recvCh

	// 模拟处理订单的耗时操作
	time.Sleep(1 * time.Second)

	// 更新库存
	product.Quantity -= 1

	fmt.Printf("Order %d processed. Updated quantity of %s: %d\n", orderID, product.Name, product.Quantity)
}

func main() {
	// 创建单向通道用于从供应商获取商品信息
	productCh := make(chan Product)

	// 启动goroutine从供应商获取商品信息
	go fetchProductInfo(productCh)

	// 模拟处理订单
	for i := 1; i <= 3; i++ {
		// 启动goroutine处理订单并更新库存
		go func() {
			processOrder(productCh, i)
		}()
	}

	// 等待一段时间，确保所有订单都被处理完毕
	time.Sleep(5 * time.Second)
}
