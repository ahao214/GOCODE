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

//直接交换法
func ReverseStr2(str string) string {
	ch := []rune(str)
	for i, j := 0, len(ch)-1; i < j; i, j = i+1, j-1 {
		ch[i] = ch[i] ^ ch[j]
		ch[j] = ch[i] ^ ch[j]
		ch[i] = ch[i] ^ ch[j]
	}
	return string(ch)

}

//方法功能：实现字符串反转
func reverse_word(ch []rune, front, end int) {
	for front < end {
		ch[front] = ch[front] ^ ch[end]
		ch[end] = ch[front] ^ ch[end]
		ch[front] = ch[front] ^ ch[end]
		front++
		end--
	}
}

//反转字符串中的单词
func swapWrods(str string) string {
	ch := []rune(str)
	l1 := len(ch)
	//对整个字符串进行字符反转操作
	reverse_word(ch, 0, l1-1)
	begin := 0
	for i := 1; i < l1; i++ {
		if ch[i] == ' ' {
			reverse_word(ch, begin, i-1)
			begin = i + 1
		}
	}
	reverse_word(ch, begin, l1-1)
	return string(ch)
}

func main() {
	str := "abcdefg"
	fmt.Println("临时变量法：")
	fmt.Printf("字符串%v翻转后为：%v", str, ReverseStr1(str))
	fmt.Println()
	fmt.Println("直接交换法")
	str2 := "abc"
	fmt.Printf("字符串%v翻转后为：%v", str2, ReverseStr2(str2))
	fmt.Println()
	fmt.Println("实现单词反转")
	strWord := "how  are you"
	fmt.Printf("字符串%v翻转之后为：%v", strWord, swapWrods(strWord))
}
