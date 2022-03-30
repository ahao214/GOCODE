package main

//496. 下一个更大元素 I
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m, n := len(nums1), len(nums2)
	res := make([]int, m)
	for i, num := range nums1 {
		j := 0
		//查找数组2中跟数组1中相等的元素
		for j < n && nums2[j] != num {
			j++
		}
		//在数组2中查找比相等元素更大的第一个元素
		k := j + 1
		for k < n && nums2[k] < nums2[j] {
			k++
		}
		if k < n {
			res[i] = nums2[k]
		} else {
			res[i] = -1
		}
	}
	return res
}
