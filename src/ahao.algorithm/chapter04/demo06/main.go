package main

import (
	"fmt"
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

func main() {
	k := 3
	arr := []int{4, 0, 1, 0, 2, 3}
	fmt.Println("第", k, "小的值为", FindSmallK(arr, 0, len(arr)-1, k))
}
