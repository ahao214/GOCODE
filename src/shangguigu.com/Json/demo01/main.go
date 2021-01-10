package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Birthday string  `json:"birthday"`
	Sal      float64 `json:"sal"`
	Skill    string  `json:"skill"`
}

//将结构体序列化
func testStruct() {
	monster := Monster{
		Name:     "牛魔王",
		Age:      500,
		Birthday: "1900-09-09",
		Sal:      10000,
		Skill:    "牛魔拳",
	}

	//将monster序列化
	data, err := json.Marshal(monster)
	if err != nil {
		fmt.Printf("序列化错误err=%v\n", err)
	}
	fmt.Printf("结构体Monster，序列化后的结果是：%v\n", string(data))
}

//将map进行序列化
func testMap() {
	//定义一个map
	var a map[string]interface{}
	//使用make
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 15
	a["skill"] = "吐火"
	a["address"] = "火云洞"

	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化错误，err=%v\n", err)
	}
	fmt.Printf("a map序列化结果是：%v\n", string(data))

}

//对切片进行序列化
func testSlice() {
	var slice []map[string]interface{}
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "Tom"
	a["age"] = 12
	a["animal"] = "猫"
	slice = append(slice, a)

	var b map[string]interface{}
	b = make(map[string]interface{})
	b["name"] = "Jerry"
	b["age"] = 10
	b["animal"] = "老鼠"
	slice = append(slice, b)

	var c map[string]interface{}
	c = make(map[string]interface{})
	c["name"] = "Jack"
	c["age"] = 14
	c["animal"] = "狗"
	slice = append(slice, c)

	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化错误：%v\n", err)
	}
	fmt.Printf("切片序列化结果是：%v\n", string(data))

}

func main() {

	//将结构体、map和切片进行序列

	//演示结构体的序列化
	testStruct()
	//演示map的序列化
	testMap()
	//演示切片的序列化
	testSlice()
}
