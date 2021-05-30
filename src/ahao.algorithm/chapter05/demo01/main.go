package main

import (
	"fmt"
	. "github.com/isdamir/gotype"
)

//求一个字符串的所有排列

//递归方法
func Permutation(str []rune, start int) {
	if str == nil {
		return
	}

	//完成全排列后输出当前排列的字符串
	if start == len(str)-1 {
		fmt.Println(string(str))
	} else {
		for i := start; i < len(str); i++ {
			//交换start与i所在位置的字符
			SwapRune(str, start, i)
			//固定第一个字符，对剩余的字符进行全排列
			Permutation(str, start+1)
			//还原start与i所在位置的字符
			SwapRune(str, start, i)
		}
	}
}

func PermutaationStr(s string) {
	Permutation([]rune(s), 0)
}

//非递归方法

//方法功能：翻转字符串
func Reverse(str []rune, begin, end int) {
	for i, j := begin, end; i < j; i, j = i+1, j-1 {
		SwapRune(str, i, j)
	}
}

//方法功能：根据当前字符串的获取下一个组合

func main() {
	fmt.Println("递归方法")
	PermutaationStr("abc")
}
