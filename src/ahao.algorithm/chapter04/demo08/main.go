package main

import (
	"fmt"
	. "github.com/isdamir/gotype"
	"math"
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

//最小距离法
func minDistanceDyn(a, b, c []int) int {
	aLen := len(a)
	bLen := len(b)
	cLen := len(c)
	minDist := math.MaxInt64

	//数组a的下标
	i := 0
	//数组b的下标
	j := 0
	//数组c的下标
	k := 0

	for true {
		curDist := Max3(Abs(a[i]-b[j]), Abs(a[i]-c[k]), Abs(b[j]-c[k]))
		if curDist < minDist {
			minDist = curDist
		}
		//找出当前遍历到三个数组中的最小值
		min := Min3(a[i], b[j], c[k])
		if min == a[i] {
			i++
			if i >= aLen {
				break
			}
		} else if min == b[j] {
			j++
			if j >= bLen {
				break
			}
		} else {
			k++
			if k >= cLen {
				break
			}
		}
	}
	return minDist
}

//数学运算法
func minDistanceMath(a, b, c []int) int {
	aLen := len(a)
	bLen := len(b)
	cLen := len(c)
	i := 0
	j := 0
	k := 0
	//最小的绝对值元和
	MinSum := Abs(a[i]-b[j]) + Abs(a[i]-c[k]) + Abs(b[j]-c[k])
	for cnt := 0; cnt <= aLen+bLen+cLen; cnt++ {
		//计算三个绝对值的和，与最小值做比较
		Sum := Abs(a[i]-b[j]) + Abs(a[i]-c[k]) + Abs(b[j]-c[k])
		MinSum = Min(MinSum, Sum)
		//找出当前遍历到三个数组中的最小值
		min := Min3(a[i], b[j], c[k])
		if min == a[i] {
			i++
			if i >= aLen {
				break
			}
		} else if min == b[j] {
			j++
			if j >= bLen {
				break
			}
		} else {
			k++
			if k >= cLen {
				break
			}
		}
	}
	return MinSum
}

func main() {
	a := []int{3, 4, 5, 7, 15}
	b := []int{10, 12, 14, 16, 17}
	c := []int{20, 21, 23, 24, 37, 30}
	fmt.Println("蛮力法")
	fmt.Println("最小距离为：", minDistance(a, b, c))
	fmt.Println("最小距离法")
	fmt.Println("最小距离为：", minDistanceDyn(a, b, c))
	fmt.Println("数学运算法")
	fmt.Println("最小距离为：", minDistanceMath(a, b, c))
}
