package main

import (
	"fmt"
	"math"
)

//找出数组中第K小的数

//方法功能：在数组arr中找出第K小的值
//输入参数：arr为整数数组，low为数组起始下标，high为数组右边界的下标，k为整数
//返回值：数组中第k小的值
func FindSmallK(arr []int, low, high, k int) int {
	i := low
	j := high
	splitElem := arr[i]
	//把小于等于splitElem的数放到数组中splitElem的左边，大于splitElem的值放到右边
	for i < j {
		for i < j && arr[j] >= splitElem {
			j--
		}
		if i < j {
			arr[i] = arr[j]
			i++
		}
		for i < j && arr[i] <= splitElem {
			i++
		}
		if i < j {
			arr[j] = arr[i]
			j--
		}
	}

	arr[i] = splitElem

	//splitElem在子数组arr[low~high]中下标的偏移量
	subArrayIndex := i - low
	//如果splitElem在arr[low~high]所在的位置恰好为k-1，那么它就是第k小的元素
	if subArrayIndex == k-1 {
		return arr[i]
	} else if subArrayIndex > k-1 {
		return FindSmallK(arr, low, i-1, k)
	} else {
		return FindSmallK(arr, i+1, high, k-(i-low)-1)
	}
}

//查找数组中前三名
func FindTopThree(arr []int) (r1, r2, r3 int) {
	if arr == nil || len(arr) < 3 {
		return
	}
	r1 = math.MinInt64
	r2 = math.MaxInt64
	r3 = math.MinInt64
	for _, v := range arr {
		if v > r1 {
			r3 = r2
			r2 = r1
			r1 = v
		} else if v > r2 && v != r1 {
			r3 = r2
			r2 = v
		} else if v > r3 && v != r2 {
			r3 = v
		}
	}
	return
}

func main() {
	k := 3
	arr := []int{4, 0, 1, 0, 2, 3}
	fmt.Println("第", k, "小的值为", FindSmallK(arr, 0, len(arr)-1, k))

	arrTop := []int{4, 7, 1, 2, 3, 5, 3, 6, 3, 2}
	r1, r2, r3 := FindTopThree(arrTop)
	fmt.Println("前三名分别为：", r1, r2, r3)
}
