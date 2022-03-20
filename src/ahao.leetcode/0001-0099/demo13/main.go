package main

import (
	"fmt"
)

//题目：13.罗马数字转整数
//说明：罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。
//字符          数值
//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000
//输入: "III"
//输出: 3

func romanToInt(s string) int {
	if s == "" {
		return 0
	}
	num, lastInt, total := 0, 0, 0
	for i := 0; i < len(s); i++ {
		char := s[len(s)-(i+1) : len(s)-i]
		num = roman[char]
		if num < lastInt {
			total = total - num
		} else {
			total = total + num
		}
		lastInt = num
	}
	return total
}

var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func main() {
	fmt.Println("罗马数字转整数")
	s := "LVIII"
	fmt.Println(romanToInt(s))
}
