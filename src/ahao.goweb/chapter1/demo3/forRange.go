package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"go":  100,
		"web": 200,
	}
	for key, value := range m {
		fmt.Println(key, value)
	}

	var str = "加油——China"
	for k, v := range str {
		fmt.Printf("key:%d value:0x%x\n", k, v)
	}

}
