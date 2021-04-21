package main

import (
	"fmt"
	"strconv"
)

//91. 解码方法
//一条包含字母 A-Z 的消息通过以下映射进行了 编码 ：
//'A' -> 1
//'B' -> 2
//...
//'Z' -> 26
//输入：s = "12"
//输出：2
//解释：它可以解码为 "AB"（1 2）或者 "L"（12）。

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[0] = 1
	if s[:1] == "0" {
		dp[1] = 0
	} else {
		dp[1] = 1
	}
	for i := 2; i <= len(s); i++ {
		lastNum, _ := strconv.Atoi(s[i-1 : i])
		if lastNum >= 1 && lastNum <= 9 {
			dp[i] += dp[i-1]
		}
		lastNum, _ = strconv.Atoi(s[i-2 : i])
		if lastNum >= 10 && lastNum <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
}

func main() {
	s := "12"
	fmt.Println(numDecodings(s))
}
