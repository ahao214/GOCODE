package main

/*
349. 两个数组的交集
给定两个数组，编写一个函数来计算它们的交集。
输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2]
*/
func intersection(nums1 []int, nums2 []int) []int {
	res := []int{}
	if nums1 == nil || len(nums1) == 0 || nums2 == nil || len(nums2) == 0 {
		return res
	}
	m := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		m[nums1[i]] = 1
	}
	for i := 0; i < len(nums2); i++ {
		if val, ok := m[nums2[i]]; ok {
			if val == 1 {
				m[nums2[i]]--
				res = append(res, nums2[i])
			}
		}
	}
	return res
}
