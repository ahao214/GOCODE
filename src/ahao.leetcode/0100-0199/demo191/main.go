package main

import (
	"fmt"
	"math/bits"
)

//题目：191. 位1的个数
//说明：编写一个函数，输入是一个无符号整数（以二进制串的形式），返回其二进制表达式中数字位数为 '1' 的个数（也被称为汉明重量）。
//输入：00000000000000000000000000001011
//输出：3
//解释：输入的二进制串 00000000000000000000000000001011 中，共有三位为 '1'。

func main() {
	res := hammingWeight(00000000000000000000000000001011)
	res1 := hammingWeight1(00000000000000000000000000001011)
	fmt.Println(res)
	fmt.Println(res1)
}

// 解法一
func hammingWeight(num uint32) int {
	return bits.OnesCount(uint(num))
}

// 解法二
func hammingWeight1(num uint32) int {
	count := 0
	for num != 0 {
		num = num & (num - 1)
		count++
	}
	return count
}
