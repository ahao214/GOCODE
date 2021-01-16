package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score float32
	Sex   string
}

func (s Monster) Print() {
	fmt.Println("---Start()---")
	fmt.Println(s)
	fmt.Println("---End()---")
}

func (s Monster) GetSum(n, m int) int {
	return n + m
}

func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func TestStruct(b interface{}) {
	typ := reflect.TypeOf(b)
	val := reflect.ValueOf(b)
	//获取b对应的类别
	kd := val.Kind()
	//判断是否是结构体
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	//获取该结构体有几个字段
	num := val.NumField()
	fmt.Printf("struct has %d fields\n", num)

	//遍历结构体的所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("field %d 的值为：%v\n", i, val.Field(i))
		//获取到struct的标签，注意需要通过reflect.Type来获取tag标签的值
		tagVal := typ.Field(i).Tag.Get("json")
		//如果该字段有tag标签就显示，否则就不显示
		if tagVal != "" {
			fmt.Printf("field %d 的tag为=%v\n", i, tagVal)
		}
	}

	//获取到该结构体有多少个方法
	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)

	//获取到第二个方法并调用它
	//方法的排序默认是按照函数名的排序，按照ASCII码
	val.Method(1).Call(nil)

	//获取到结构体的第一个方法Method(0)
	var params []reflect.Value
	params = append(params, reflect.ValueOf(20))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params)
	fmt.Println("res=", res[0].Int())

}

func main() {
	var a Monster = Monster{
		Name:  "黄鼠狼精",
		Age:   100,
		Score: 200,
		Sex:   "male",
	}
	TestStruct(a)
}
