package main

import "math/bits"

/*
762. 二进制表示中质数个计算置位
给你两个整数 left 和 right ，在闭区间 [left, right] 范围内，统计并返回 计算置位位数为质数 的整数个数。

计算置位位数 就是二进制表示中 1 的个数。

例如， 21 的二进制表示 10101 有 3 个计算置位。
*/
func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func countPrimeSetBits(left, right int) (ans int) {
	for x := left; x <= right; x++ {
		if isPrime(bits.OnesCount(uint(x))) {
			ans++
		}
	}
	return
}

//判断质数优化
func countPrimeSetBits2(left, right int) (ans int) {
	for x := left; x <= right; x++ {
		if 1<<bits.OnesCount(uint(x))&665772 != 0 {
			ans++
		}
	}
	return
}
