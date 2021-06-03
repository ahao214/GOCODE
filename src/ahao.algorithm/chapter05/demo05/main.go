package main

import (
	"fmt"
)

//判断两个字符串的包含关系
//直接法
func isContain(str1, str2 string) bool {
	len1 := len(str1)
	len2 := len(str2)
	if len1 < len2 {
		for i := 0; i < len1; i++ {
			for j := 0; j < len2; j++ {
				if str1[i] == str2[j] {
					break
				}
				if j >= len2 {
					return false
				}
			}
		}
	} else {
		for i := 0; i < len2; i++ {
			for j := 0; j < len1; j++ {
				if str1[j] == str2[i] {
					break
				}
				if j >= len1 {
					return false
				}
			}
		}
	}
	return true
}

func main() {
	str1 := "abcdef"
	str2 := "acf"
	fmt.Println("直接法：")
	fmt.Printf("%v与%v", str1, str2)
	if isContain(str1, str2) {
		fmt.Println("有包含关系")
	} else {
		fmt.Println("没有包含关系")
	}
}
