package main

/**
 * 三数之中的最大值
 * @param num1: An integer
 * @param num2: An integer
 * @param num3: An integer
 * @return: an interger
 */
func maxOfThreeNumbers(num1 int, num2 int, num3 int) int {
	return max(num1, max(num2, num3))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
