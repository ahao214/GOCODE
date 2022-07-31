package main

/*
1331. 数组序号转换
*/

func arrayRankTransform(arr []int) []int {
	a := append([]int{}, arr...)
	sort.Ints(a)
	ranks := map[int]int{}
	for _, v := range a {
		if _, ok := ranks[v]; !ok {
			ranks[v] = len(ranks) + 1
		}
	}
	for i, v := range arr {
		arr[i] = ranks[v]
	}
	return arr
}
