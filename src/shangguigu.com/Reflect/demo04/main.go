package main

import (
	"fmt"
	"reflect"
)

//通过反射修改
//num int的值
//修改Student的值

func refTest(b interface{}) {
	rValue := reflect.ValueOf(b)
	fmt.Println("rValue:", rValue)
	fmt.Printf("int类型的rValue的kind是：%v\n", rValue.Kind())

	rValue.Elem().SetInt(20)
}

func refString(b interface{}) {
	rValue := reflect.ValueOf(b)
	fmt.Println("rValue:", rValue)
	fmt.Printf("string类型的rValue的kind是：%v\n", rValue.Kind())

	rValue.Elem().SetString("hello go")
}

func main() {
	var num int = 10
	refTest(&num)
	fmt.Println("num的值是：", num)

	var strHello string = "hello world"
	refString(&strHello)
	fmt.Println("strHello的新的值是：", strHello)
}
