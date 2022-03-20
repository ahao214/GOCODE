package main

/*
2044. 统计按位或能得到最大值的子集数目
难度：中等
给你一个整数数组 nums ，请你找出 nums 子集 按位或 可能得到的 最大值 ，并返回按位或能得到最大值的 不同非空子集的数目 。

如果数组 a 可以由数组 b 删除一些元素（或不删除）得到，则认为数组 a 是数组 b 的一个 子集 。如果选中的元素下标位置不一样，则认为两个子集 不同 。

对数组 a 执行 按位或 ，结果等于 a[0] OR a[1] OR ... OR a[a.length - 1]（下标从 0 开始）。
*/
func countMaxOrSubsets(nums []int) int {
	n, m := len(nums), 0
	for _, num := range nums {
		m |= num
	}

	var dfs func(idx, val int) int
	dfs = func(idx, val int) int {
		if idx == n {
			if val == m {
				return 1
			}
			return 0
		}
		if nVal := val | nums[idx]; nVal == m {
			return dfs(idx+1, val) + (1 << (n - 1 - idx))
		} else {
			return dfs(idx+1, val) + dfs(idx+1, nVal)
		}
	}

	return dfs(0, 0)
}

//位运算
func countMaxOrSubsets2(nums []int) (ans int) {
	maxOr := 0
	for i := 1; i < 1<<len(nums); i++ {
		or := 0
		for j, num := range nums {
			if i>>j&1 == 1 {
				or |= num
			}
		}
		if or > maxOr {
			maxOr = or
			ans = 1
		} else if or == maxOr {
			ans++
		}
	}
	return
}

//回溯
func countMaxOrSubsets3(nums []int) (ans int) {
	maxOr := 0
	var dfs func(int, int)
	dfs = func(pos, or int) {
		if pos == len(nums) {
			if or > maxOr {
				maxOr = or
				ans = 1
			} else if or == maxOr {
				ans++
			}
			return
		}
		dfs(pos+1, or|nums[pos])
		dfs(pos+1, or)
	}
	dfs(0, 0)
	return
}
