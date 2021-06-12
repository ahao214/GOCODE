package main

import (
	"fmt"
)

//判断一个字符串是否包含重复字符

//蛮力法

func isDup(str string) bool {
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] == str[j] {
				return true
			}
		}
	}
	return false
}

func main() {
	str := "GOOD"
	fmt.Println("方法一：")
	if isDup(str) {
		fmt.Println(str, "中有重复字符")
	} else {
		fmt.Println(str, "中没有重复字符")
	}

}
