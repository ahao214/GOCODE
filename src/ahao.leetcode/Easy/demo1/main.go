package main

//两数之和
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
