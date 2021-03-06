package main

import (
	"fmt"
)

//寻找最多的覆盖点
func maxCover(a []int, L int) int {
	count := 2
	maxCount := 1 //最长覆盖的点数
	start := 0    //覆盖坐标的真实位置
	n := len(a)
	for i, j := 0, 1; i < n && j < n; {
		for (j < n) && (a[j]-a[i] <= L) {
			j++
			count++
		}
		j--
		count--
		if count > maxCount {
			start = i
			maxCount = count
		}
		i++
		j++
	}
	fmt.Println("覆盖的坐标点：")
	for i := start; i < start+maxCount; i++ {
		fmt.Print(a[i], " ")
	}
	fmt.Println()
	return maxCount
}

func main() {
	arr := []int{1, 3, 7, 8, 10, 11, 12, 13, 15, 16, 17, 19, 25}
	fmt.Println("最长覆盖点数：", maxCover(arr, 8))

}
