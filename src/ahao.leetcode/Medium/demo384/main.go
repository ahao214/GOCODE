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

//Fisher-Yates洗牌算法
type Solution2 struct {
	nums2, original2 []int
}

func Constructor2(nums2 []int) Solution2 {
	return Solution2{nums2, append([]int(nil), nums2...)}
}

func (s *Solution2) Reset2() []int {
	copy(s.nums2, s.original2)
	return s.nums2
}

func (s *Solution2) Shuffle2() []int {
	n := len(s.nums2)
	for i := range s.nums2 {
		j := i + rand.Intn(n-i)
		s.nums2[i], s.nums2[j] = s.nums2[j], s.nums2[i]
	}
	return s.nums2
}
