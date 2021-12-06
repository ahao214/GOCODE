package main

//201. 数字范围按位与
func rangeBitwiseAnd(left int, right int) int {
	for left < right {
		right = right & (right - 1)
	}
	return right
}

func rangeBitwiseAnd2(left int, right int) int {
	shift := 0
	for left < right {
		left, right = left>>1, right>>1
		shift++
	}
	return right << shift
}
