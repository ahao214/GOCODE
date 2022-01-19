package main

import "sort"

/*
350. 两个数组的交集 II
给你两个整数数组 nums1 和 nums2 ，请你以数组形式返回两数组的交集。
返回结果中每个元素出现的次数，应与元素在两个数组中都出现的次数一致
（如果出现次数不一致，则考虑取较小值）。
可以不考虑输出结果的顺序。
*/
func intersect(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	var res []int
	for _, n := range nums1 {
		m[n]++
	}
	for _, n := range nums2 {
		if m[n] > 0 {
			res = append(res, n)
			m[n]--
		}
	}
	return res
}

func intersect1(nums1 []int, nums2 []int) []int {
	res := []int{}
	if nums1 == nil || nums2 == nil || len(nums1) == 0 || len(nums2) == 0 {
		return res
	}
	m := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		m[nums1[i]]++
	}
	for i := 0; i < len(nums2); i++ {
		if val, ok := m[nums2[i]]; ok {
			if val > 0 {
				res = append(res, nums2[i])
			}
			m[nums2[i]]--
		}
	}
	return res
}

//两个已经排序的数组
func intersect2(nums1 []int, nums2 []int) []int {
	res := []int{}
	if nums1 == nil || nums2 == nil || len(nums1) == 0 || len(nums2) == 0 {
		return res
	}

	sort.Ints(nums1)
	sort.Ints(nums2)
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			res = append(res, nums1[i])
			i++
			j++
		}
		if i < len(nums1) && j < len(nums2) && nums1[i] < nums2[j] {
			i++
		}
		if i < len(nums1) && j < len(nums2) && nums2[j] < nums1[i] {
			j++
		}
	}
	return res
}
