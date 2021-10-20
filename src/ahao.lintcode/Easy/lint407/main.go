package main

//407 加一
func plusOne(digits []int) []int {
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
