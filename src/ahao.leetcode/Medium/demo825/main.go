package main

import "sort"

//825. 适龄的朋友

//方法一：排序+双指针
func numFriendRequests(ages []int) (ans int) {
	sort.Ints(ages)
	left, right := 0, 0
	for _, age := range ages {
		if age < 15 {
			continue
		}
		for ages[left]*2 <= age+14 {
			left++
		}
		for right+1 < len(ages) && ages[right+1] <= age {
			right++
		}
		ans += right - left
	}
	return
}

//方法二：计数排序 + 前缀和
func numFriendRequests1(ages []int) (ans int) {
	const mx = 121
	var cnt, pre [mx]int
	for _, age := range ages {
		cnt[age]++
	}
	for i := 1; i < mx; i++ {
		pre[i] = pre[i-1] + cnt[i]
	}
	for i := 15; i < mx; i++ {
		if cnt[i] > 0 {
			bound := i/2 + 8
			ans += cnt[i] * (pre[i] - pre[bound-1] - 1)
		}
	}
	return
}
