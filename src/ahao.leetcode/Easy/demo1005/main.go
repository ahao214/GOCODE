package main

import "sort"

//1005. K 次取反后最大化的数组和
func largestSumAfterKNegations(nums []int, k int) int {
	ans := 0
	dic := map[int]int{}
	for _, n := range nums {
		dic[n]++
		ans += n
	}

	for i := -100; i < 0 && k != 0; i++ {
		if dic[i] > 0 {
			ops := min(k, dic[i])
			ans -= i * ops * 2
			dic[-i] += ops
			k -= ops
		}
	}
	if k > 0 && k%2 == 1 && dic[0] == 0 {
		for i := 1; i <= 100; i++ {
			if dic[i] > 0 {
				ans -= i * 2
				break
			}
		}
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func largestSumAfterKNegations2(A []int, K int) int {
	sort.Ints(A)
	minIdx := 0
	for i := 0; i < K; i++ {
		A[minIdx] = -A[minIdx]
		if A[minIdx+1] < A[minIdx] {
			minIdx++
		}
	}
	sum := 0
	for _, a := range A {
		sum += a
	}
	return sum
}
