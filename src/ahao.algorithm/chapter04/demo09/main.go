package main

import (
	"fmt"
	. "github.com/isdamir/gotype"
	"math"
)

//求数组中绝对值最小的数

//顺序比较法
func FindMin(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	min := math.MaxInt64
	for _, v := range arr {
		if Abs(v) < Abs(min) {
			min = v
		}
	}
	return min
}

//数学性质法
func FindMinMath(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	begin := 0
	mid := 0
	end := len(arr) - 1
	//数组中没有负数
	if arr[0] > 0 {
		return arr[end]
	}
	if arr[end] <= 0 {
		return arr[end]
	}
	absMin := 0

	//数组中既有正数也有负数
	for true {
		mid = begin + (end+begin)/2
		//如果等于2，那么就是绝对值最小的数
		if arr[mid] == 0 {
			return 0
		} else if arr[mid] > 0 {
			//如果大于0，正负数的分界点在左侧
			//继续在数组的左半部分查找
			if arr[mid-1] > 0 {
				end = mid - 1
			} else if arr[mid-1] == 0 {
				return 0
			} else {
				//找到正负数的分界点
				break
			}
		} else {
			//在数组有半部分继续查找
			if arr[mid+1] < 0 {
				begin = mid + 1
			} else if arr[mid+1] == 0 {
				return 0
			} else {
				//找到正负数的分界点
				break
			}
		}
	}

	//获取正负数分界点出绝对值最小的值
	if arr[mid] > 0 {
		if arr[mid] < Abs(arr[mid-1]) {
			absMin = arr[mid]
		} else {
			absMin = arr[mid-1]
		}
	} else {
		if Abs(arr[mid]) < arr[mid+1] {
			absMin = arr[mid]
		} else {
			absMin = arr[mid+1]
		}
	}
	return absMin
}

func main() {
	arr := []int{-10, -4, -2, 7, 15, 50}
	fmt.Println("顺序比较法：")
	fmt.Println("绝对值最小的数为：", FindMin(arr))

	fmt.Println("数学性质发")
	fmt.Println("绝对值最小的数为：", FindMinMath(arr))
}
