package main

import "fmt"

//题目：7. 整数反转
//说明：给你一个 32 位的有符号整数 x ，返回 x 中每位上的数字反转后的结果
//输入：x = 123
//输出：321
//解释：
func reverse(x int) int {
	tmp := 0
	for x != 0 {
		tmp = tmp*10 + x%10
		x = x / 10
	}
	if tmp > 1<<31-1 || tmp < -(1<<31) {
		return 0
	}
	return tmp
}

func main() {
	n := 123
	res := reverse(n)
	fmt.Println(res)
}
