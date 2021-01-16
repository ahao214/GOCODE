package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

//对结构体的反射
func reflectTest(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Println("rType:", rType)

	rValue := reflect.ValueOf(b)
	fmt.Println("rValue:", rValue)

	rKindone := rValue.Kind()
	rKindtwo := rType.Kind()
	fmt.Printf("rKindone:%v,rKindtwo:%v\n", rKindone, rKindtwo)

	//将rValue转成interface{}
	iv := rValue.Interface()
	fmt.Printf("iv=%v iv's type=%T\n", iv, iv)

	//将interface{}通过断言转成需要的类型
	stu, ok := iv.(Student)
	if ok {
		fmt.Printf("学生的姓名：%s,学生的年龄是：%d\n", stu.Name, stu.Age)
	}

}

func main() {
	//定义一个Student的实例
	stu := Student{Name: "Tom", Age: 20}
	reflectTest(stu)
}
