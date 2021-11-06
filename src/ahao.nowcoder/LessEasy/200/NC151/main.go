package main

/**
 * 求出a、b的最大公约数。
 * @param a int
 * @param b int
 * @return int
 */
func gcd(a int, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}
