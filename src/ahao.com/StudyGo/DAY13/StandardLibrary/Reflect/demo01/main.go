package main

import (
	"fmt"
	"reflect"
)

//反射获取interface类型信息
func reflect_type(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Println("类型是：", t)
	//kind()可以获取具体类型
	k := t.Kind()
	fmt.Println(k)
	switch k {
	case reflect.Float64:
		fmt.Printf("a is float64\n")
	case reflect.String:
		fmt.Println("string")
	}
}

//空接口与反射
func main() {
	var x float64 = 3.4
	reflect_type(x)
}
