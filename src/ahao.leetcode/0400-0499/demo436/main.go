package main

import "sort"

/*
436. 寻找右区间
给你一个区间数组 intervals ，其中 intervals[i] = [starti, endi] ，且每个 starti 都 不同 。

区间 i 的 右侧区间 可以记作区间 j ，并满足 startj >= endi ，且 startj 最小化 。

返回一个由每个区间 i 的 右侧区间 在 intervals 中对应下标组成的数组。如果某个区间 i 不存在对应的 右侧区间 ，则下标 i 处的值设为 -1 。


示例 1：

输入：intervals = [[1,2]]
输出：[-1]
解释：集合中只有一个区间，所以输出-1。
示例 2：

输入：intervals = [[3,4],[2,3],[1,2]]
输出：[-1,0,1]
解释：对于 [3,4] ，没有满足条件的“右侧”区间。
对于 [2,3] ，区间[3,4]具有最小的“右”起点;
对于 [1,2] ，区间[2,3]具有最小的“右”起点。
示例 3：

输入：intervals = [[1,4],[2,3],[3,4]]
输出：[-1,2,-1]
解释：对于区间 [1,4] 和 [3,4] ，没有满足条件的“右侧”区间。
对于 [2,3] ，区间 [3,4] 有最小的“右”起点。


提示：

1 <= intervals.length <= 2 * 104
intervals[i].length == 2
-106 <= starti <= endi <= 106
每个间隔的起点都 不相同
*/

//二分查找
func findRightInterval(intervals [][]int) []int {
	for i := range intervals {
		intervals[i] = append(intervals[i], i)
	}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	n := len(intervals)
	ans := make([]int, n)
	for _, p := range intervals {
		i := sort.Search(n, func(i int) bool { return intervals[i][0] >= p[1] })
		if i < n {
			ans[p[2]] = intervals[i][2]
		} else {
			ans[p[2]] = -1
		}
	}
	return ans
}

//双指针
func findRightInterval2(intervals [][]int) []int {
	n := len(intervals)
	type pair struct{ x, i int }
	starts := make([]pair, n)
	ends := make([]pair, n)
	for i, p := range intervals {
		starts[i] = pair{p[0], i}
		ends[i] = pair{p[1], i}
	}
	sort.Slice(starts, func(i, j int) bool { return starts[i].x < starts[j].x })
	sort.Slice(ends, func(i, j int) bool { return ends[i].x < ends[j].x })

	ans := make([]int, n)
	j := 0
	for _, p := range ends {
		for j < n && starts[j].x < p.x {
			j++
		}
		if j < n {
			ans[p.i] = starts[j].i
		} else {
			ans[p.i] = -1
		}
	}
	return ans
}
