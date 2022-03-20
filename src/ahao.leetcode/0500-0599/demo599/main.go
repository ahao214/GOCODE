package main

import "math"

/*
599. 两个列表的最小索引总和
难度：简单
假设 Andy 和 Doris 想在晚餐时选择一家餐厅，并且他们都有一个表示最喜爱餐厅的列表，每个餐厅的名字用字符串表示。

你需要帮助他们用最少的索引和找出他们共同喜爱的餐厅。 如果答案不止一个，则输出所有答案并且不考虑顺序。 你可以假设答案总是存在。
*/
func findRestaurant(list1, list2 []string) (ans []string) {
	index := make(map[string]int, len(list1))
	for i, s := range list1 {
		index[s] = i
	}
	indexSum := math.MaxInt32
	for i, s := range list2 {
		if j, ok := index[s]; ok {
			if i+j < indexSum {
				indexSum = i + j
				ans = []string{s}
			} else if i+j == indexSum {
				ans = append(ans, s)
			}
		}
	}
	return
}
