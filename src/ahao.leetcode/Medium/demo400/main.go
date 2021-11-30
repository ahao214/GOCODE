package main

import "math"

//400. 第 N 位数字

//直接计算
func findNthDigit(n int) int {
	if n < 10 {
		return n
	}
	d := 1
	for count := 9; n > d*count; count *= 10 {
		n -= d * count
		d++
	}

	index := n - 1
	start := int(math.Pow10(d - 1))
	num := start + index/d
	digitIndex := index % d
	return num / int(math.Pow10(d-digitIndex)) % 10
}
