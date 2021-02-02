package main

import (
	"fmt"
)

func main() {
	//创建一个映射，存储颜色以及颜色对应的十六进制代码
	colors := map[string]string{
		"AliccBlue":   "#f0f8ff",
		"Coral":       "#ff7fff",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}
	//显示映射里的所有颜色
	for key, value := range colors {
		fmt.Printf("key:%s value:%s\n", key, value)
	}

	//调用函数来移除指定的键
	RemoveColor(colors, "Coral")

	fmt.Println()
	//显示映射里的所有颜色
	for key, value := range colors {
		fmt.Printf("key:%s value:%s\n", key, value)
	}
}

func RemoveColor(colors map[string]string, key string) {
	delete(colors, key)

}
