package main

/*
57. 插入区间
给你一个 无重叠的 ，按照区间起始端点排序的区间列表。
在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。
输入：intervals = [[1,3],[6,9]], newInterval = [2,5]
输出：[[1,5],[6,9]]
*/
func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)
	if intervals == nil || len(intervals) == 0 {
		res = append(res, newInterval)
		return res
	}

	start, end := newInterval[0], newInterval[1]
	left, right := make([][]int, 0), make([][]int, 0)
	for i := 0; i < len(intervals); i++ {
		curStart, curEnd := intervals[i][0], intervals[i][1]
		if curEnd < start {
			left = append(left, intervals[i])
		} else if curStart > end {
			right = append(right, intervals[i])
		} else {
			start = min(start, curStart)
			end = max(end, curEnd)
		}
	}
	tmp := []int{start, end}
	res = append(res, left...)
	res = append(res, tmp)
	res = append(res, right...)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
