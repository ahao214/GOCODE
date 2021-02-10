package main

import (
	"fmt"
)

type StockPosition struct {
	ticket     string
	sharePrice float32
	count      float32
}

//获取StockPosition的值(价格)
func (s StockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type Car struct {
	make  string
	model string
	price float32
}

//获取Car的值(价格)
func (c Car) getValue() float32 {
	return c.price
}

//定义具有价值的不同事务的合同
type valueable interface {
	getValue() float32
}

func showValue(asset valueable) {
	fmt.Printf("资产的价值是：%f\n", asset.getValue())
}

func main() {
	var o valueable = StockPosition{"Good", 577.00, 20.4}
	showValue(o)
	o = Car{"BMW", "M3", 665000}
	showValue(o)
}
