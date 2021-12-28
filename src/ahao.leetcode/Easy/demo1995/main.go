package main

//1995. 统计特殊四元组
func countQuadruplets(nums []int) (ans int) {
	for a, x := range nums {
		for b := a + 1; b < len(nums); b++ {
			for c := b + 1; c < len(nums); c++ {
				for _, y := range nums[c+1:] {
					if x+nums[b]+nums[c] == y {
						ans++
					}
				}
			}
		}
	}
	return
}

//使用哈希表存储nums[d]
func countQuadruplets1(nums []int) (ans int) {
	cnt := map[int]int{}
	for c := len(nums) - 2; c >= 2; c-- {
		cnt[nums[c+1]]++
		for a, x := range nums[:c] {
			for _, y := range nums[a+1 : c] {
				if sum := x + y + nums[c]; cnt[sum] > 0 {
					ans += cnt[sum]
				}
			}
		}
	}
	return
}
