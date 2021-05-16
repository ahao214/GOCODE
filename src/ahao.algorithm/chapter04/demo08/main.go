package main

import (
	"fmt"
	. "github.com/isdamir/gotype"
)

//求解最小三元组距离

//蛮力法
func minDistance(a, b, c []int) int {
	minDist := Max3(Abs(a[0]-b[0]), Abs(a[0]-c[0]), Abs(b[0]-c[0]))
	dist := 0
	for _, aData := range a {
		for _, bData := range b {
			for _, cData := range c {
				//求距离
				dist = Max3(Abs(aData-bData), Abs(aData-cData), Abs(bData-cData))
				//找出最小距离
				if minDist > dist {
					minDist = dist
				}
			}
		}
	}
	return minDist
}

func main() {
	a := []int{3, 4, 5, 7, 15}
	b := []int{10, 12, 14, 16, 17}
	c := []int{20, 21, 23, 24, 37, 30}
	fmt.Println("蛮力法")
	fmt.Println("最小距离为：", minDistance(a, b, c))
}
