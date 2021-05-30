package main

import (
	"bytes"
	"fmt"
)

//求两个字符串的最长公共子串

//动态规划法
//方法功能：获取两个字符串的最长公共字串
func getMaxSubStr(str1, str2 string) string {
	len1 := len(str1)
	len2 := len(str2)
	var buf bytes.Buffer
	maxI := 0
	max := 0
	//申请新的空间来记录公共字串长度信息
	M := make([][]int, len1+1)
	for i, _ := range M {
		M[i] = make([]int, len2+1)
	}
	//通过利用递归公式填写新建的二维数组(公共字串的长度信息)
	for i := 1; i < len1+1; i++ {
		for j := 1; j < len2+1; j++ {
			if str1[i-1] == str2[j-1] {
				M[i][j] = M[i-1][j-1] + 1
				if M[i][j] > max {
					max = M[i][j]
					maxI = i
				}
			} else {
				M[i][j] = 0
			}
		}
	}
	for i := maxI - max; i < maxI; i++ {
		buf.WriteByte(str1[i])
	}
	return buf.String()
}

//滑动比较法
func getMaxSub2(str1, str2 string) string {
	len1 := len(str1)
	len2 := len(str2)
	var buf bytes.Buffer
	maxLen := 0
	maxLenEnd1 := 0
	for i := 0; i < len1+len2; i++ {
		s1begin := 0
		s2begin := 0
		tmpMaxLne := 0
		if i < len1 {
			s1begin = len1 - i
		} else {
			s2begin = i - len1
		}
		j := 0
		for j = 0; (s1begin+j < len1) && (s2begin+j < len2); j++ {
			if str1[s1begin+j] == str2[s2begin+j] {
				tmpMaxLne++
			} else {
				if tmpMaxLne > maxLen {
					maxLen = tmpMaxLne
					maxLenEnd1 = s1begin + j
				} else {
					tmpMaxLne = 0
				}
			}
		}
		if tmpMaxLne > maxLen {
			maxLen = tmpMaxLne
			maxLenEnd1 = s1begin + j
		}
	}
	for i := maxLenEnd1 - maxLen; i < maxLenEnd1; i++ {
		buf.WriteByte(str1[i])
	}
	return buf.String()
}

func main() {
	str1 := "abccade"
	str2 := "dgcadde"
	fmt.Println("动态规划方法")
	fmt.Println(getMaxSubStr(str1, str2))

	fmt.Println("滑动比较法")
	fmt.Println(getMaxSub2(str1, str2))
}
