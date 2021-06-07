package main

import (
	"fmt"
	. "github.com/isdamir/gotype"
)

//判断字符串是否是整数

//递归方法
var flag bool

func strToIntSub(arr []rune, length int) int {
	if length > 1 {
		if !IsNumber(arr[len(arr)-1]) {
			fmt.Println("不是数字")
			flag = false
			return -1
		}
		if arr[0] == '-' {
			return strToIntSub(arr, length-1)*10 - int(arr[length-1]-'0')
		} else {
			return strToIntSub(arr, length-1)*10 + int(arr[length-1]-'0')
		}
	} else {
		if arr[0] == '-' {
			return 0
		} else {
			if !IsNumber(arr[0]) {
				fmt.Println("不是数字")
				flag = false
				return -1
			}
			return int(arr[0] - '0')
		}
	}
}

func strToInt(s string) int {
	arr := []rune(s)
	flag = true
	if arr[0] == '+' {
		return strToIntSub(arr[1:], len(arr)-1)
	} else {
		return strToIntSub(arr, len(arr))
	}
}

func printResult(s string) {
	re := strToInt(s)
	if flag {
		fmt.Println(re)
	}
}

func main() {
	fmt.Println("递归方法：")
	printResult("-543")
	printResult("543")
	printResult("+543")
	printResult("++43")
}
