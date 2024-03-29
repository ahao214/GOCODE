package main

/**
 * NC57 反转数字
 * @param x int整型
 * @return int整型
 */
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
