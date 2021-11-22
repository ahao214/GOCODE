package main

//986. 区间列表的交集
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	res := [][]int{}
	if firstList == nil || len(firstList) == 0 || secondList == nil || len(secondList) == 0 {
		return res
	}
	i := 0
	j := 0
	firstLen := len(firstList)
	secondLen := len(secondList)
	for i < firstLen && j < secondLen {
		start1 := firstList[i][0]
		end1 := firstList[i][1]
		start2 := secondList[j][0]
		end2 := secondList[j][1]
		//选择开始时间晚的
		start := max(start1, start2)
		////选择结束时间早的
		end := min(end1, end2)
		if start <= end {
			res = append(res, []int{start, end})
		}
		if end1 <= end2 {
			i++
		} else {
			j++
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
