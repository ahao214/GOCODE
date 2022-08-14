package main

/*
1413. 逐步求和得到正数的最小值
*/

//贪心
func minStartValue(nums []int) int {
	accSum, accSumMin := 0, 0
	for _, num := range nums {
		accSum += num
		accSumMin = min(accSumMin, accSum)
	}
	return -accSumMin + 1
}



//二分查找方法
func minStartValue2(nums []int) int {
	m := nums[0]
	for _, num := range nums[1:] {
		m = min(m, num)
	}
	return 1 + sort.Search(-m*len(nums), func(startValue int) bool {
		startValue++
		for _, num := range nums {
			startValue += num
			if startValue <= 0 {
				return false
			}
		}
		return true
	})
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
