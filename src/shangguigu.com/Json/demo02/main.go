package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name    string
	Age     int
	Address string
}

//将json字符串反序列化成结构体
func unStruct() {
	str := "{\"Name\":\"牛魔王\",\"Age\":200,\"Address\":\"芭蕉洞\"}"
	var monster Monster
	//反序列化
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("反序列化失败：%v\n", err)
	}
	fmt.Printf("反序列化后monster=%v\n", monster)

}

//将json字符串，反序列化成map
func unMap() {
	str := "{\"Name\":\"牛魔王\",\"Age\":200,\"Address\":\"芭蕉洞\"}"
	var a map[string]interface{}
	//反序列化
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("反序列化失败：%v\n", err)
	}
	fmt.Printf("反序列化后a=%v\n", a)

}

//将json字符串，反序列化成切片
func unSlice() {
	str := "[{\"Name\":\"牛魔王\",\"Age\":200,\"Address\":\"芭蕉洞\"},{\"Name\":\"青牛精\",\"Age\":100,\"Address\":\"青牛山\"}]"
	var slice []map[string]interface{}

	//反序列化
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("反序列化失败：%v\n", err)
	}
	fmt.Printf("反序列化后slice=%v\n", slice)

}

func main() {
	//反序列化成结构体
	unStruct()
	//反序列化map
	unMap()
	//反序列化切片
	unSlice()

}
