package main

import (
	"fmt"
)

//题目：第650题：只有两个键的键盘
//说明：
//输入: 3
//输出: 3
//解释: 最初, 我们只有一个字符 'A'。
//第 1 步, 我们使用 Copy All 操作。
//第 2 步, 我们使用 Paste 操作来获得 'AA'。
//第 3 步, 我们使用 Paste 操作来获得 'AAA'
func main() {
	n := 30
	fmt.Println(minSteps(n))
}

func minSteps(n int) int {
	res := 0
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			res += i
			n /= i
		}
	}
	return res
}
