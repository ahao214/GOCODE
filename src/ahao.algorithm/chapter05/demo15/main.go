package main

import (
	"fmt"
)

//求最长递增子序列的长度
//动态规划法
func getMaxAscendingLen(str string) int {
	maxLen := make([]int, len(str))
	maxAcsendingLen := 1
	for i, _ := range str {
		maxLen[i] = 1
		for j := 0; j < i; j++ {
			if str[j] < str[i] && maxLen[j] > maxLen[i]-1 {
				maxLen[i] = maxLen[j] + 1
				maxAcsendingLen = maxLen[i]
			}
		}
	}
	return maxAcsendingLen
}

func main() {
	fmt.Println("最长递增子序列的长度为：", getMaxAscendingLen("xbcdza"))
}
