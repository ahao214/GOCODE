package main

/*
66.加一
给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
你可以假设除了整数 0 之外，这个整数不会以零开头。
*/

func plusOne(digits []int) []int {
	var result []int
	addone := 0
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += addone
		addone = 0
		if i == len(digits)-1 {
			digits[i]++
		}
		if digits[i] == 10 {
			addone = 1
			digits[i] = digits[i] % 10
		}
	}
	if addone == 1 {
		result = make([]int, 1)
		result[0] = 1
		result = append(result, digits...) //将digits切片中的元素追加到result里面
	} else {
		result = digits
	}
	return result
}

func plusOne1(digits []int) []int {
	res := make([]int, 0)
	if digits == nil || len(digits) == 0 {
		return res
	}
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + carry
		carry = sum / 10
		digits[i] = sum % 10
	}
	if carry == 1 {
		res = []int{1}
		digits = append(res, digits...)
	}
	return digits
}
