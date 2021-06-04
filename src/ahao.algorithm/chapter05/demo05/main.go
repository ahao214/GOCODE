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

//空间换时间法
func isContain2(str1, str2 string) bool {
	len1 := len(str1)
	len2 := len(str2)
	//用来记录52个字母的出现情况
	flag := make([]int, 52)
	//记录段字符串中不同字符出现的个数
	count := 0
	//分别用来记录较短和较长的字符串
	var shortStr string
	var longStr string
	//分别用来记录较长和较短字符的长度
	var minLen int
	var maxLen int
	if len1 < len2 {
		shortStr = str1
		minLen = len1
		longStr = str2
		maxLen = len2
	} else {
		shortStr = str2
		minLen = len2
		longStr = str1
		maxLen = len1
	}
	var k rune
	//遍历短字符串
	for i := 0; i < minLen; i++ {
		if shortStr[i] >= 'A' && shortStr[i] <= 'Z' {
			k = rune(shortStr[i]) - 'A'
		} else {
			k = rune(shortStr[i]) - 'a' + 26
		}
		if flag[k] == 0 {
			flag[k] = 1
			count++
		}
	}

	//遍历长字符串
	for i := 0; i < maxLen; i++ {
		if longStr[i] >= 'A' && longStr[i] <= 'Z' {
			k = rune(longStr[i]) - 'A'
		} else {
			k = rune(longStr[i]) - 'a' + 26
		}
		if flag[k] == 1 {
			flag[k] = 0
			count--
			if count == 0 {
				return true
			}
		}
	}
	return false
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

	fmt.Println("空间换时间法：")
	if isContain2(str1, str2) {
		fmt.Println("有包含关系")
	} else {
		fmt.Println("没有包含关系")
	}
}
