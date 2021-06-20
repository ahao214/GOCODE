package main

import (
	"fmt"
	. "github.com/isdamir/gotype"
)

//参数replaceWight用来表示替换操作与插入删除操作相比的倍数
func edit(s1, s2 string, replaceWight int) int {
	if s1 == "" || s2 == "" {
		return 0
	}
	if s1 == "" {
		return len(s2)
	}
	if s2 == "" {
		return len(s1)
	}
	len1, len2 := len(s1), len(s2)
	//申请二维数组来存储中间的计算结果
	D := make([][]int, len1+1)
	for i, _ := range D {
		D[i] = make([]int, len2+1)
	}
	for i := 0; i < len1+1; i++ {
		D[i][0] = i
	}
	for i := 0; i < len2+1; i++ {
		D[0][i] = i
	}
	for i := 1; i < len1+1; i++ {
		for j := 1; j < len2+1; j++ {
			if s1[i-1] == s2[j-1] {
				D[i][j] = Min3(D[i-1][j], D[i][j-1]+1, D[i-1][j-1])
			} else {
				D[i][j] = Min3(D[i-1][j]+1, D[i][j-1]+1, D[i-1][j-1]+replaceWight)
			}
		}
	}
	fmt.Println("--------------------------")
	for i := 0; i < len1+1; i++ {
		for j := 0; j < len2+1; j++ {
			fmt.Print(D[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println("--------------------------")
	return D[len1][len2]
}

//求字符串的编辑距离
func main() {
	s1 := "bciln"
	s2 := "fciling"
	fmt.Println("第一问：")
	fmt.Println("编辑距离：", edit(s1, s2, 1))
	fmt.Println("第二问：")
	fmt.Println("编辑距离：", edit(s1, s2, 2))
}
