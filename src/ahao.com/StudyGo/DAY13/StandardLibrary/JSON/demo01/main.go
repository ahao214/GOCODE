package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string
	Hobby string
}

//结构体生成json
func main() {
	p := Person{"51.com", "女"}
	//编码json
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
	//格式化输出
	b, err = json.MarshalIndent(p, "", "    ")
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}
