package main

import (
	"encoding/json"
	"fmt"
)

//通过map生成json
func main() {
	student := make(map[string]interface{})
	student["name"] = "枯藤"
	student["age"] = 18
	student["sex"] = "男"
	b, err := json.Marshal(student)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}
