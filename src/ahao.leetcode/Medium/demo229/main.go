package main

//229. 求众数 II

func majorityElement(nums []int) []int {
	res := []int{}
	dic := map[int]int{}
	for _, v := range nums {
		dic[v]++
	}
	for k, v := range dic {
		if v > len(nums)/3 {
			res = append(res, k)
		}
	}
	return res
}
