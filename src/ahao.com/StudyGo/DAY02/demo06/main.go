package main

import (
	"fmt"
)

//Map的使用
//map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用

func main() {

	userInfo := map[string]string{
		"username": "ahao",
		"pwd":      "123",
	}
	fmt.Println(userInfo)

	score := make(map[string]int, 8)
	score["张三"] = 100
	score["李四"] = 200
	fmt.Println(score)
	fmt.Println(score["李四"])
	fmt.Printf("type of a:%T\n", score)

	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	v, ok := score["张三"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}

	//map的遍历
	for k, v := range score {
		fmt.Println(k, v)
	}

	//使用delete()函数删除键值对
	delete(score, "李四")
	fmt.Println("删除之后的数据")
	for k, v := range score {
		fmt.Println(k, v)
	}
}
