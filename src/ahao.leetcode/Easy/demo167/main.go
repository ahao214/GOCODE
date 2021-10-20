package main

//167. 两数之和 II - 输入有序数组
func twoSum(numbers []int, target int) []int {
	res := []int{-1, -1}
	dic := make(map[int]int)
	for i, n := range numbers {
		tmp := target - n
		if _, ok := dic[tmp]; ok {
			if i > dic[tmp] {
				res[0] = dic[tmp] + 1
				res[1] = i + 1
			} else {
				res[0] = i + 1
				res[1] = dic[tmp] + 1
			}
			return res
		}
		dic[n] = i
	}
	return res
}
