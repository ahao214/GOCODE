package main

import (
	"sort"
)

//1079. 活字印刷

/*
你有一套活字字模 tiles，其中每个字模上都刻有一个字母 tiles[i]。返回你可以印出的非空字母序列的数目。

注意：本题中，每个活字字模只能使用一次。

输入："AAB"
输出：8
解释：可能的序列为 "A", "B", "AA", "AB", "BA", "AAB", "ABA", "BAA"。
*/
func numTilePossibilities(tiles string) int {
	if len(tiles) <= 1 {
		return len(tiles)
	}
	res := 0
	chars := []byte(tiles)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	cur := []byte{}
	visited := make([]bool, len(tiles))
	backtrack(chars, cur, visited, &res)
	return res
}

func backtrack(chars []byte, cur []byte, visited []bool, res *int) {
	if len(cur) > 0 {
		(*res)++
	}
	for i := 0; i < len(chars); i++ {
		if visited[i] {
			continue
		}
		if i > 0 && chars[i] == chars[i-1] && visited[i-1] == false {
			continue
		}
		cur = append(cur, chars[i])
		visited[i] = true
		backtrack(chars, cur, visited, res)
		visited[i] = false
		cur = cur[:len(cur)-1]
	}
}
