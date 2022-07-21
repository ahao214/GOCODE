package main

/*
556. 下一个更大元素 III
*/

//方法一：下一个排列
func nextGreaterElement(n int) int {
	nums := []byte(strconv.Itoa(n))
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i < 0 {
		return -1
	}

	j := len(nums) - 1
	for j >= 0 && nums[i] >= nums[j] {
		j--
	}
	nums[i], nums[j] = nums[j], nums[i]
	reverse(nums[i+1:])
	ans, _ := strconv.Atoi(string(nums))
	if ans > math.MaxInt32 {
		return -1
	}
	return ans
}

func reverse(a []byte) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}


//方法二：数学
