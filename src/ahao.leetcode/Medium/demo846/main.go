package main

import "sort"

//846. 一手顺子
/*
Alice 手中有一把牌，她想要重新排列这些牌，分成若干组，使每一组的牌数都是 groupSize ，
并且由 groupSize 张连续的牌组成。
给你一个整数数组 hand 其中 hand[i] 是写在第 i 张牌，和一个整数 groupSize 。
如果她可能重新排列这些牌，返回 true ；否则，返回 false 。
*/
func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}

	sort.Ints(hand)
	cnt := map[int]int{}
	for _, num := range hand {
		cnt[num]++
	}

	for _, x := range hand {
		if cnt[x] == 0 {
			continue
		}
		for num := x; num < x+groupSize; num++ {
			if cnt[num] == 0 {
				return false
			}
			cnt[num]--
		}
	}
	return true
}
