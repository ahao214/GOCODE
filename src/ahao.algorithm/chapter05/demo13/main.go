package main

import (
	"fmt"
	"sort"
)

//找到由其他单词组成的最长单词
type LongestWord struct {
}

//判断字符串str是否在字符串数组中
func (p *LongestWord) find(strArr []string, str string) bool {
	for _, v := range strArr {
		if str == v {
			return true
		}
	}
	return false
}

//方法功能；判断字符串word是否能有数组strArray中的其他单词组成
//参数：word为待判断的后缀子串，length为待判断字符串的长度
func (p *LongestWord) isContain(strArr []string, word string, length int) bool {
	ll := len(word)
	//递归的结束条件，当字符串长度为0时，说明字符串已经遍历完了
	if ll == 0 {
		return true
	}

	//循环取字符串的所有前缀
	for i := 1; i <= ll; i++ {
		//取到的子串为自己
		if i == length {
			return false
		}
		str := string(word[0:i])
		if p.find(strArr, str) {
			//查找玩字符串的前缀后，递归判断后面的子串能否由其他单词组成
			if p.isContain(strArr, string(word[i:]), length) {
				return true
			}
		}
	}
	return false
}

//找出能由数组中其他字符串组成的最长字符串
func (p *LongestWord) GetLongestStr(strArr []string) string {
	//对字符串由大到小排序
	sort.Slice(strArr, func(i, j int) bool {
		return len(strArr[i]) > len(strArr[j])
	})
	//贪心地从最长的字符串开始判断
	for _, v := range strArr {
		if p.isContain(strArr, v, len(v)) {
			return v
		}
	}
	return ""
}

func main() {
	strArr := []string{"test", "tester", "testertest", "testing", "apple", "seattle", "banana", "batting", "ngcat", "batti", "bat", "testingtester", "testbattingcat"}
	lw := new(LongestWord)
	longestStr := lw.GetLongestStr(strArr)
	if longestStr == "" {
		fmt.Println("不存在这样的字符串")
	} else {
		fmt.Println("最长的字符串为：", longestStr)
	}

}
