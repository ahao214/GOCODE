package main

/**
 * 简单计算器
 * (正数不包括0)
 * @param a: An integer
 * @param operator: A character, +, -, *, /.
 * @param b: An integer
 * @return: The result
 */
func calculate(a int, operator byte, b int) int {
	res := 0
	if operator == '+' {
		res = a + b
	}
	if operator == '-' {
		res = a - b
	}
	if operator == '*' {
		res = a * b
	}
	if operator == '/' {
		res = a / b
	}

	return res

}
