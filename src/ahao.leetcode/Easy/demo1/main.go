package main

//1.两数之和
//在数组中找到 2 个数之和等于给定值的数字，结果返回 2 个数字在数组中的下标
//输入：nums = [2, 7, 11, 15], target = 9,
//解释：nums[0] + nums[1] = 2 + 7 = 9,
//输出：return [0, 1]
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}
	return nil
}

func twoNum2(nums []int, target int) []int {
	res := []int{-1, -1}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res[0] = i
				res[1] = j
				return res
			}
		}
	}
	return res
}

func twoNums3(nums []int, target int) []int {
	res := []int{-1, -1}
	lookUp := make(map[int]int)
	for i, n := range nums {
		tmp := target - n
		if _, ok := lookUp[tmp]; ok {
			res[0] = lookUp[tmp]
			res[1] = i
			return res
		}
		lookUp[n] = i
	}
	return res
}
