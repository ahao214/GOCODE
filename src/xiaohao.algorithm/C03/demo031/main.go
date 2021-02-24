package main

import (
	"fmt"
)

//第137题：只出现一次的数字Ⅱ
func main() {
	nums := []int{2, 2, 3, 2}
	res := singleNumber(nums)
	fmt.Println(res)
}

func singleNumber(nums []int) int {
	number, res := 0, 0
	for i := 0; i < 64; i++ {
		//初始化每一位1的个数为0
		number = 0
		for _, k := range nums {
			//通过右移i位的方式，计算每一位1的个数
			number += (k >> i) & 1
		}
		//最终将抵消后剩余的1放到对应的位数上
		res |= (number) % 3 << i
	}
	return res
}
