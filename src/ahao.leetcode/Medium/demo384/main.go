package main

import "math/rand"

//384. 打乱数组

//暴力方法
type Solution struct {
	nums, original []int
}

func Constructor(nums []int) Solution {
	return Solution{nums, append([]int(nil), nums...)}
}

func (s *Solution) Reset() []int {
	copy(s.nums, s.original)
	return s.nums
}

func (s *Solution) Shuffle() []int {
	shuffled := make([]int, len(s.nums))
	for i := range shuffled {
		j := rand.Intn(len(s.nums))
		shuffled[i] = s.nums[j]
		s.nums = append(s.nums[:j], s.nums[j+1:]...)
	}
	s.nums = shuffled
	return s.nums
}
