package main

/*
565. 数组嵌套
*/

//图
func arrayNesting(nums []int) (ans int) {
	vis := make([]bool, len(nums))
	for i := range vis {
		cnt := 0
		for !vis[i] {
			vis[i] = true
			i = nums[i]
			cnt++
		}
		if cnt > ans {
			ans = cnt
		}
	}
	return
}

//原地标记
func arrayNesting2(nums []int) (ans int) {
	n := len(nums)
	for i := range nums {
		cnt := 0
		for nums[i] < n {
			i, nums[i] = nums[i], n
			cnt++
		}
		if cnt > ans {
			ans = cnt
		}
	}
	return
}
