package main

/**
 * 771 · 二阶阶乘
 * @param n: the given number
 * @return:  the double factorial of the number
 */
func doubleFactorial(n int) int64 {
	res := 1
	for i := n; i > 0; i -= 2 {
		res *= i
	}
	return int64(res)
}
