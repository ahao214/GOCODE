package main

import "sort"

//594. 最长和谐子序列

//枚举
func findLHS(nums []int) int {
	ans := 0
	sort.Ints(nums)
	begin := 0
	for end, n := range nums {
		if n-nums[begin] > 1 {
			begin++
		}
		if n-nums[begin] == 1 && end-begin+1 > ans {
			ans = end - begin + 1
		}
	}
	return ans
}

func findLHS2(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	res := make(map[int]int, len(nums))
	for _, num := range nums {
		if _, exist := res[num]; exist {
			res[num]++
			continue
		}
		res[num] = 1
	}
	longest := 0
	for k, c := range res {
		if n, exist := res[k+1]; exist {
			if c+n > longest {
				longest = c + n
			}
		}
	}
	return longest
}

//哈希表
func findLHS3(nums []int) (ans int) {
	cnt := map[int]int{}
	for _, num := range nums {
		cnt[num]++
	}
	for n, c := range cnt {
		if c1 := cnt[n+1]; c1 > 0 && c+c1 > ans {
			ans = c + c1
		}
	}
	return ans
}
