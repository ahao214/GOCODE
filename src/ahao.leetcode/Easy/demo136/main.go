package main

//136. 只出现一次的数字
func singleNumber(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	dic := make(map[int]int)
	for _, v := range nums {
		dic[v]++
	}
	for k, v := range dic {
		if v == 1 {
			return k
		}
	}
	return -1
}

func singleNumberXor(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}
