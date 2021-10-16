package main

//1207. 独一无二的出现次数
func uniqueOccurrences(arr []int) bool {
	if arr == nil || len(arr) == 0 {
		return true
	}
	dic := make(map[int]int)
	for _, v := range arr {
		dic[v]++
	}
	countDic := make(map[int]int)
	for _, v := range dic {
		_, ok := countDic[v]
		if !ok {
			countDic[v]++
		} else {
			return false
		}
	}
	return true
}
