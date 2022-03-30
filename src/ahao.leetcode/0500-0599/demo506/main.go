package main

import (
	"sort"
	"strconv"
)

//506. 相对名次
var desc = [3]string{"Gold Medal", "Silver Medal", "Bronze Medal"}

func findRelativeRanks(score []int) []string {
	n := len(score)
	type pair struct {
		score, idx int
	}
	arr := make([]pair, n)
	for i, s := range score {
		arr[i] = pair{s, i}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].score > arr[j].score
	})
	res := make([]string, n)
	for i, p := range arr {
		if i < 3 {
			res[p.idx] = desc[i]
		} else {
			res[p.idx] = strconv.Itoa(i)
		}
	}
	return res
}
