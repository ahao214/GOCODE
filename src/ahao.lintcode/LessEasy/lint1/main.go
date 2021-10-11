package main

//1 · A + B 问题
/**
 * @param a: An integer
 * @param b: An integer
 * @return: The sum of a and b
 */
func aplusb(a int, b int) int {
	if a == b {
		return a * 2
	}
	return (a*a - b*b) / (a - b)
}
