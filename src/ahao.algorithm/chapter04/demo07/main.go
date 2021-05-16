package main

import (
	"fmt"
	. "github.com/isdamir/gotype" //引入定义的数据结构
	"math"
)

//求数组中两个元素的最小距离

//蛮力法
func minDistance(arr []int, num1, num2 int) int {
	if arr == nil || len(arr) <= 0 {
		return math.MaxInt64
	}
	minDis := math.MaxInt64
	dist := 0
	for i, v := range arr {
		if v == num1 {
			for j, v := range arr {
				if v == num2 {
					dist = Abs(i - j) //当前遍历的num1和num2的距离
					if dist < minDis {
						minDis = dist
					}
				}
			}
		}
	}
	return minDis
}

//动态规划
func minDynDistance(arr []int, num1, num2 int) int {
	if arr == nil || len(arr) <= 0 {
		return math.MaxInt64
	}
	//上次遍历到num1的位置
	lastPos1 := -1

	//上次遍历到num2的位置
	lastPos2 := -1

	//num1和num2的最小距离
	minDis := math.MaxInt64

	for i, v := range arr {
		if v == num1 {
			lastPos1 = i
			if lastPos2 >= 0 {
				minDis = Min(minDis, lastPos1-lastPos2)
			}
		}
		if v == num2 {
			lastPos2 = i
			if lastPos1 >= 0 {
				minDis = Min(minDis, lastPos2-lastPos1)
			}
		}
	}
	return minDis
}

func main() {
	arr := []int{4, 5, 6, 4, 7, 4, 6, 4, 7, 8, 5, 6, 4, 3, 10, 8}
	num1 := 4
	num2 := 8
	fmt.Println("蛮力法")
	fmt.Println(minDistance(arr, num1, num2))

	fmt.Println("动态规划法")
	fmt.Println(minDynDistance(arr, num1, num2))
}
