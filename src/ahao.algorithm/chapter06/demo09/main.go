package main

import (
	"bytes"
	"fmt"
)

func intToBinary(n int) string {
	if n < 0 {
		fmt.Println("不支持负数")
		return ""
	}
	hexNum := 8 * 8
	var pBinary bytes.Buffer
	for i := 0; i < hexNum; i++ {
		tmpBinary := (n << uint(i)) >> (uint(hexNum) - 1)
		if tmpBinary == 0 {
			pBinary.WriteRune('0')
		} else {
			pBinary.WriteRune('1')
		}
	}
	return pBinary.String()
}

func intToHex(s int) string {
	hex := ""
	remainder := 0
	for s != 0 {
		remainder = s % 16
		if remainder < 10 {
			hex = string(remainder+'0') + hex
		} else {
			hex = string(remainder-10+'A') + hex
		}
		s = s >> 4
	}
	return hex
}

//如何把十进制数(long型)分别以二进制和十六进制形式输出
func main() {
	fmt.Println("10的二进制输出为：", intToBinary(10))
	fmt.Println("10的十六进制输出为：", intToHex(10))
}
