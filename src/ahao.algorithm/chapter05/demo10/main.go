package main

import (
	"fmt"
)

//求字符串里的最长回文子串

//动态规划法
type Tests struct {
	startIndex int
	len        int
}

func (p *Tests) getLongestPalindrome(str string) {
	ll := len(str)
	if ll < 1 {
		return
	}
	p.startIndex = 0
	p.len = 1
	//申请额外的存储空间记录查找的历史信息
	historyRecord := make([][]int, ll)
	for i, _ := range historyRecord {
		historyRecord[i] = make([]int, ll)
	}
	//初始化长度为1的回文字符串信息
	for i := 0; i < ll; i++ {
		historyRecord[i][i] = 1
	}
	//初始化长度为2的回文字符串信息
	for i := 0; i < ll-1; i++ {
		if str[i] == str[i+1] {
			historyRecord[i][i+1] = 1
			p.startIndex = i
			p.len = 2
		}
	}
	//查找从长度为3开始的回文字符串
	for pLen := 3; pLen <= ll; pLen++ {
		for i := 0; i < ll-pLen+1; i++ {
			var tmp = i + pLen - 1
			if str[i] == str[tmp] && historyRecord[i+1][tmp-1] == 1 {
				historyRecord[i][tmp] = 1
				p.startIndex = i
				p.len = pLen
			}
		}
	}
}

//对字符串str，以c1和c2为中心向两侧扩展寻找回文子串
func (p *Tests) expandBothSide(str string, c1, c2 int) {
	n := len(str)
	for c1 >= 0 && c2 < n && str[c1] == str[c2] {
		c1--
		c2++
	}
	tmpStartIndex := c1 + 1
	tmpLen := c2 - c1 - 1
	if tmpLen > p.len {
		p.len = tmpLen
		p.startIndex = tmpStartIndex
	}
}

//中心扩展法
//对字符串str，以c1和c2为中心向两侧扩展寻找回文子串
func (p *Tests) expendBothSide2(str string, c1, c2 int) {
	n := len(str)
	for c1 >= 0 && c2 < n && str[c1] == str[c2] {
		c1--
		c2++
	}
	tmpStartIndex := c1 + 1
	tmpLen := c2 - c1 - 1
	if tmpLen > p.len {
		p.len = tmpLen
		p.startIndex = tmpStartIndex
	}
}

//找出字符串最长的回文子串
func (p *Tests) getLongestPalindrome2(str string) {
	n := len(str)
	if n < 1 {
		return
	}
	for i, _ := range str {
		//找回文字符串长度为奇数的情况
		p.expendBothSide2(str, i, i)
		//找回文字符串长度为偶数的情况
		p.expendBothSide2(str, i, i+1)
	}
}

func main() {
	str := "abcdefgfedxyz"
	t := &Tests{}
	fmt.Println("方法一：")
	t.getLongestPalindrome(str)
	if t.startIndex != -1 && t.len != -1 {
		fmt.Println("最长的回文子串为：", string(str[t.startIndex:t.startIndex+t.len]))
	} else {
		fmt.Println("查找失败")
	}

	fmt.Println("方法二：")
	t2 := Tests{}
	t2.getLongestPalindrome2(str)
	if t2.startIndex != -1 && t2.len != -1 {
		fmt.Println("最长的回文子串为：", string(str[t2.startIndex:t2.startIndex+t.len]))
	} else {
		fmt.Println("查找失败")
	}
}
