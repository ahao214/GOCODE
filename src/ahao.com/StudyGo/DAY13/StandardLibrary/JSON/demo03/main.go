package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Age       int    `json:"age,string"`
	Name      string `json:"name"`
	Niubility string `json:"niubility"`
}

//解析到结构体
func main() {
	//假数据
	b := []byte(`{"age":"18","name":"枯藤","marry":false}`)
	var p Person
	err := json.Unmarshal(b, &p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
}
