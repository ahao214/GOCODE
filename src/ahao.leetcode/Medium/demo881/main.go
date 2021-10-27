package main

import "sort"

//881. 救生艇
func numRescueBoats(people []int, limit int) int {
	//排序+左右指针
	if people == nil || len(people) == 0 {
		return 0
	}
	res := 0
	left, right := 0, len(people)-1
	sort.Ints(people)
	for left <= right {
		sum := people[left] + people[right]
		if left == right {
			sum = people[left]
		}
		if sum <= limit {
			left++
		}
		right--
		res++
	}
	return res
}
