package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
	}
	//修改
	m["a"] = 10
	//新增
	m["c"] = 30

	n := make(map[string]int)
	for i := 0; i < 8; i++ {
		n[string('a'+i)] = i
	}

	for i := 0; i < 4; i++ {
		for k, v := range n {
			fmt.Print(k, ":", v, " ")
		}
		fmt.Println()
	}

	k := make(map[int]int)
	for i := 0; i < 10; i++ {
		k[i] = i + 10
	}
	for a := range k {
		if a == 5 {
			k[100] = 1000
		}
		delete(k, a)
		fmt.Println(a, k)
	}

}
