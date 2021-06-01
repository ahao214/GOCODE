package main

import (
	"fmt"
)

//对字符串进行翻转

//临时变量方法
func ReverseStr1(str string) string {
	ch := []rune(str)
	for i, j := 0, len(ch)-1; i < j; i, j = i+1, j-1 {
		tmp := ch[i]
		ch[i] = ch[j]
		ch[j] = tmp
	}
	return string(ch)

}

func main() {
	str := "abcdefg"
	fmt.Println("临时变量法：")
	fmt.Printf("字符串%v翻转后为：%v", str, ReverseStr1(str))
	fmt.Println()
}
