package main

import (
	"fmt"
	. "github.com/isdamir/gotype"
	"sort"
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
func getNextPermutation(str []rune) bool {
	end := len(str) - 1
	cur := end
	suc := 0
	for cur != 0 {
		//从后向前开始遍历字符串
		suc = cur
		cur--
		if str[cur] < str[suc] {
			//相邻递增的字符,cur指向较小的字符
			//找出cur后面最小字符tmp
			tmp := end
			for str[tmp] < str[cur] {
				tmp--
			}
			//交换cur与tmp
			SwapRune(str, cur, tmp)
			//把cur后面的字符串进行翻转
			Reverse(str, suc, end)
			return true
		}
	}
	return false
}

func PermutationMath(s string) {
	data := []rune(s)
	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	}) //升序排列字符数组
	for true {
		fmt.Println(string(data))
		if !getNextPermutation(data) {
			break
		}
	}
}

func main() {
	fmt.Println("递归方法")
	PermutaationStr("abc")
	fmt.Println("非递归方法")
	PermutationMath("abc")
}
