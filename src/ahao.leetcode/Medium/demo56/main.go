package main

import (
	"sort"
)

/*
56. 合并区间
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
*/
func merge(intervals [][]int) [][]int {
	res := make([][]int, 0)
	if intervals == nil || len(intervals) == 0 {
		return res
	}

	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		curStart := intervals[i][0]
		curEnd := intervals[i][1]

		preEnd := intervals[i-1][1]
		if preEnd >= curStart {
			intervals[i-1][1] = max(preEnd, curEnd)
			intervals[i] = intervals[i-1]
		} else {
			res = append(res, intervals[i])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
