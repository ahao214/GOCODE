package main

/*
905. 按奇偶排序数组
*/

//两次遍历
func sortArrayByParity(nums []int) []int {
	ans := make([]int, 0, len(nums))
	for _, num := range nums {
		if num%2 == 0 {
			ans = append(ans, num)
		}
	}
	for _, num := range nums {
		if num%2 == 1 {
			ans = append(ans, num)
		}
	}
	return ans
}

//双指针+一次遍历
func sortArrayByParity2(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	left, right := 0, n-1
	for _, num := range nums {
		if num%2 == 0 {
			ans[left] = num
			left++
		} else {
			ans[right] = num
			right--
		}
	}
	return ans
}

//原地交换
func sortArrayByParity3(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		for left < right && nums[left]%2 == 0 {
			left++
		}
		for left < right && nums[right]%2 == 1 {
			right--
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	return nums
}
