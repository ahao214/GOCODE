package main

import (
	"fmt"
	"strconv"
)

//获取x和y异或的结果
func xor1(x, y int) int {
	BITS := strconv.IntSize //以具体平台为例
	res := 0
	var xoredBit int
	for i := BITS - 1; i >= 0; i-- {
		//获取x与y当前的bit值
		b1 := (x & (1 << uint(i))) > 0
		b2 := (y & (1 << uint(i))) > 0

		//只有这两位都是1或0的时候结果为0
		if b1 == b2 {
			xoredBit = 0
		} else {
			xoredBit = 1
		}
		res <<= 1
		res |= xoredBit
	}
	return res
}

func xor(x, y int) int {
	return (x | y) & (^x | ^y)
}

//如何不使用^操作实现异或运算
func main() {
	fmt.Println("方法一：")
	fmt.Println(xor1(3, 5))

	fmt.Println("方法二：")
	fmt.Println(xor(3, 5))
}
