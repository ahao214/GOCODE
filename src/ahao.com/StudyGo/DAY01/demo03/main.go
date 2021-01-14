package main

import (
	"fmt"
	"sort"
)

//映射(map)-建立事务关联的容器

func main() {
	sence := make(map[string]int)
	sence["one"] = 1
	fmt.Println(sence["one"])
	v := sence["two"]
	fmt.Println(v)

	m := map[string]string{
		"A": "Forward",
		"L": "Left",
		"R": "Right",
		"D": "Behind",
	}
	for k, v := range m {
		fmt.Println(k, v)
	}

	area := make(map[string]int)
	area["china"] = 960
	area["usa"] = 900
	area["russ"] = 980

	//声明一个切片报错map数据
	var arealist []string

	//变量map，并赋值到切片中
	for k, _ := range area {
		arealist = append(arealist, k)
	}
	//对切片进行排序
	sort.Strings(arealist)
	fmt.Println(arealist)

	//从map中删除键值对
	area1 := make(map[string]int)
	area1["china"] = 960
	area1["usa"] = 900
	area1["russ"] = 980
	delete(area1, "russ")
	for k, v := range area1 {
		fmt.Println(k, v)
	}

	//声明map
	var m1 map[string]string
	//使用make函数创建一个非nil的map，nil map不能赋值
	m1 = make(map[string]string)
	//给声明的map赋值
	m1["a"] = "aa"
	m1["b"] = "bb"

	//直接创建
	m2 := make(map[string]string)
	m2["c"] = "cc"

	//初始化+赋值一体化
	m3 := map[string]string{
		"d": "dd",
		"e": "ee",
	}

	//查找键值是否存在
	if v, ok := m1["a"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("key not found")
	}

	//遍历map
	for k, v := range m3 {
		fmt.Println(k, v)
	}

	for k, v := range m1 {
		fmt.Println(k, v)
	}
}
