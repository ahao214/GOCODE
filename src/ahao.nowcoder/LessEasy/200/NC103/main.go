package main

import "fmt"

/**
 * 反转字符串
 * @param str string字符串
 * @return string字符串
 */
func solve(str string) string {
	strBytes := []byte(str)
	left := 0
	right := len(strBytes) - 1
	for left < right {
		strBytes[left], strBytes[right] = strBytes[right], strBytes[left]
		left++
		right--
	}
	return string(strBytes)
}

func main() {
	str := "abc"
	fmt.Println(solve(str))
}
