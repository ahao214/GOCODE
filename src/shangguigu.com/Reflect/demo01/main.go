package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	//通过反射，获取到传入的变量的type kind 值

	//获取到type
	rType := reflect.TypeOf(b)
	fmt.Println("rType:", rType)

	//获取到value
	rValue := reflect.ValueOf(b)
	fmt.Println("rValue:", rValue)

	n := 10 + rValue.Int()
	fmt.Printf("n的值是：%d\n", n)

	//将rValue转成interface{}
	iV := rValue.Interface()
	//将interface{}通过断言转成需要的类型
	num := iV.(int)
	fmt.Println("num的值是：", num)

}

func main() {

	//演示对(基本数据类型，interface{}，reflect.Value)进行反射的基本操作

	//1.先定义一个int类型
	var num int = 100

	reflectTest01(num)

}
