package main

/**
 * JZ65 不用加减乘除做加法
 * @param num1 int整型
 * @param num2 int整型
 * @return int整型
 */
func Add(num1 int, num2 int) int {
	if num1 == 0 && num2 == 0 {
		return 0
	} else if num1 == 0 {
		return num2
	} else if num2 == 0 {
		return num1
	}
	if num1 > num2 {
		return (num1*num1 - num2*num2) / (num1 - num2)
	}
	return (num2*num2 - num1*num1) / (num2 - num1)
}
