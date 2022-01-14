package main

import "sort"

/*
373. 查找和最小的K对数字
给定两个以升序排列的整数数组 nums1 和 nums2 , 以及一个整数 k 。
定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2 。
请找到和最小的 k 个数对 (u1,v1),  (u2,v2)  ...  (uk,vk) 。
*/
func kSmallestPairs(nums1, nums2 []int, k int) (ans [][]int) {
	m, n := len(nums1), len(nums2)

	// 二分查找第 k 小的数对和
	left, right := nums1[0]+nums2[0], nums1[m-1]+nums2[n-1]+1
	pairSum := left + sort.Search(right-left, func(sum int) bool {
		sum += left
		cnt := 0
		i, j := 0, n-1
		for i < m && j >= 0 {
			if nums1[i]+nums2[j] > sum {
				j--
			} else {
				cnt += j + 1
				i++
			}
		}
		return cnt >= k
	})

	// 找数对和小于 pairSum 的数对
	i := n - 1
	for _, num1 := range nums1 {
		for i >= 0 && num1+nums2[i] >= pairSum {
			i--
		}
		for _, num2 := range nums2[:i+1] {
			ans = append(ans, []int{num1, num2})
			if len(ans) == k {
				return
			}
		}
	}

	// 找数对和等于 pairSum 的数对
	i = n - 1
	for _, num1 := range nums1 {
		for i >= 0 && num1+nums2[i] > pairSum {
			i--
		}
		for j := i; j >= 0 && num1+nums2[j] == pairSum; j-- {
			ans = append(ans, []int{num1, nums2[j]})
			if len(ans) == k {
				return
			}
		}
	}
	return
}
