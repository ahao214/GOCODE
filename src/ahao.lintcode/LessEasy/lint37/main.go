package main

import (
	"fmt"
)

/**
 * 反转一个三位整数
 * @param number: A 3-digit number.
 * @return: Reversed number.
 */
func reverseInteger(number int) int {
	tmp := 0
	for number != 0 {
		tmp = tmp*10 + number%10
		number = number / 10
	}
	if tmp > 1<<31-1 || tmp < -(1<<31) {
		return 0
	}
	return tmp
}

func main() {
	fmt.Println(reverseInteger(90))
}
