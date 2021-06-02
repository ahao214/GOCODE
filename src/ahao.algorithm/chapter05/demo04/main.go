package main

import (
	"fmt"
)

//判断两个字符串是否为换位字符串
//换位字符串是指组成字符串的字符相同

func compare(s1, s2 string) bool {
	result := true
	bCount := make([]int, 256)
	for _, v := range s1 {
		bCount[v]++
	}
	for _, v := range s2 {
		bCount[v]--
	}
	for _, v := range bCount {
		if v != 0 {
			return false
		}
	}
	return result
}

func main() {
	str1 := "aaaabbc"
	str2 := "abcbaaa"
	fmt.Printf("%v和%v", str1, str2)
	if compare(str1, str2) {
		fmt.Println("是换位字符串")
	} else {
		fmt.Println("不是换位字符串")
	}

}
