package main

import (
	"fmt"
)

//移位法
func countOne(n int) int {
	count := 0
	for n > 0 {
		if (n & 1) == 1 { //判断最后一位是否为1
			count++
		}
		n >>= 1 //移位丢掉最后一位
	}
	return count
}

//求二进制数中1的个数
func main() {
	fmt.Println("移位法：")
	fmt.Println(countOne(7))
	fmt.Println(countOne(8))
}
