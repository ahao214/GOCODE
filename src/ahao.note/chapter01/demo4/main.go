package main

import "fmt"

//字典类型
func main() {
	//创建字典类型对象
	m := make(map[string]int)
	m["a"] = 1
	x, ok := m["b"]
	fmt.Println(x, ok)
	delete(m, "a")
}
