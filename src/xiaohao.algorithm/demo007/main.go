package main

import (
	"fmt"
)

//66.加一
func plueOne(digits []int) []int {
	var result []int
	addone := 0
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += addone
		addone = 0
		if i == len(digits)-1 {
			digits[i]++
		}
		if digits[i] == 10 {
			addone = 1
			digits[i] = digits[i] % 10
		}
	}
	if addone == 1 {
		result = make([]int, 1)
		result[0] = 1
		result = append(result, digits...) //将digits切片中的元素追加到result里面
	} else {
		result = digits
	}
	return result
}

func main() {

	digits := []int{1, 2, 3, 4, 5}
	one := []int{9, 9, 9}
	two := []int{1, 2, 3, 9}
	fmt.Println(plueOne(digits))
	fmt.Println(plueOne(one))
	fmt.Println(plueOne(two))
	fmt.Println("hello world")
}
