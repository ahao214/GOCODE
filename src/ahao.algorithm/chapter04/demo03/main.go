package main

import (
	"fmt"
)

//找出旋转数组的最小元素
func getMinPara(arr []int, low, high int) int {
	//如果选择个数是0，即没有旋转，单独处理，直接返回数组头元素
	if high < low {
		return arr[0]
	}
	//只剩下一个元素一定是最小值
	if high == low {
		return arr[low]
	}
	mid := low + ((high - low) >> 1)

	//判断arr[mid]为最小值
	if arr[mid] < arr[mid-1] {
		return arr[mid]
	} else if arr[mid+1] < arr[mid] {
		//判断是否arr[mid+1]为最小值
		return arr[mid+1]
	} else if arr[high] > arr[mid] {
		//最小值一定在数组左半部分
		return getMinPara(arr, low, mid-1)
	} else if arr[mid] > arr[low] {
		//最小值一定在数组有半部分
		return getMinPara(arr, mid+1, high)
	} else {
		left := getMinPara(arr, low, mid-1)
		right := getMinPara(arr, mid+1, high)
		if left < right {
			return left
		} else {
			return right
		}
	}
}

func getMin(arr []int) int {
	if arr == nil {
		return -1
	} else {
		return getMinPara(arr, 0, len(arr)-1)
	}
}

func swap(arr []int, low, high int) {
	//交换数组low到high的内容
	for ; low < high; low, high = low+1, high-1 {
		tmp := arr[low]
		arr[low] = arr[high]
		arr[high] = tmp
	}
}

func main() {
	fmt.Println("找出最小值")
	arr := []int{5, 6, 1, 2, 3, 4}
	fmt.Println(getMin(arr))
	arr = []int{1, 1, 0, 1}
	fmt.Println(getMin(arr))
}
