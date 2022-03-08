package main

/*
798. 得分最高的最小轮调
给定一个数组 A，我们可以将它按一个非负整数 K 进行轮调，这样可以使数组变为 A[K], A[K+1], A{K+2], ... A[A.length - 1], A[0], A[1], ..., A[K-1] 的形式。此后，任何值小于或等于其索引的项都可以记作一分。

例如，如果数组为 [2, 4, 1, 3, 0]，我们按 K = 2 进行轮调后，它将变成 [1, 3, 0, 2, 4]。这将记作 3 分，因为 1 > 0 [no points], 3 > 1 [no points], 0 <= 2 [one point], 2 <= 3 [one point], 4 <= 4 [one point]。

在所有可能的轮调中，返回我们所能得到的最高分数对应的轮调索引 K。如果有多个答案，返回满足条件的最小的索引 K。
*/
func bestRotation(nums []int) int {
	n := len(nums)
	diffs := make([]int, n)
	for i, num := range nums {
		low := (i + 1) % n
		high := (i - num + n + 1) % n
		diffs[low]++
		diffs[high]--
		if low >= high {
			diffs[0]++
		}
	}
	score, maxScore, idx := 0, 0, 0
	for i, diff := range diffs {
		score += diff
		if score > maxScore {
			maxScore, idx = score, i
		}
	}
	return idx
}
