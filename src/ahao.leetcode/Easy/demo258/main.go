package main

import "fmt"

//524. 通过删除字母匹配到字典里最长单词
func addDigits(num int) int {
	for num > 0 {
		cur := 0
		for num != 0 {
			cur += num % 10
			num /= 10
		}
		num = cur
	}
	return num
}

func main() {
	fmt.Println(addDigits(38))
}
