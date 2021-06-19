package main

import (
	"fmt"
	"strings"
)

func isSubstring(str1, str2 string) bool {
	return strings.Index(str1, str2) >= 0
}

//判断str12是否可以通过旋转str1得到
func rotateSame(str1, str2 string) bool {
	len1, len2 := len(str1), len(str2)
	//判断两个字符串长度是否相等，如果不相等，不可能通过旋转得到
	if len1 != len2 {
		return false
	}
	//申请临时空间存储s1，多申请了一个空间存储
	tmp := make([]rune, 2*len1+1)
	for i, v := range str1 {
		tmp[i] = v
		tmp[i+len1] = v
	}
	tmp[2*len1] = ' '
	//判断str2是否为tmp的子串
	return isSubstring(string(tmp), str2)
}

//判断一个字符串是否由另外一个字符串旋转得到
func main() {
	str1 := "waterbottle"
	str2 := "erbottlewat"
	result := rotateSame(str1, str2)
	if result {
		fmt.Println(str2, "可以通过旋转", str1, "得到")
	} else {
		fmt.Println(str2, "不可以通过旋转", str1, "得到")
	}

}
