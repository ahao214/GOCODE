package main

import (
	"fmt"
)

//快速排序
//left 数组左边的下标
//right 数组右边的下标
func QuickSort(left int, right int, arr *[6]int) {
	l := left
	r := right
	//pivot 是中轴
	pivot := arr[(left+right)/2]

	//循环的目标是：将比pivot小的数放在左边，比pivot大的数放在右边
	for l < r {
		//从pivot的左边找到大于等于pivot的值
		for arr[l] < pivot {
			l++
		}
		//从pivot的右边找到小于等于pivot的值
		for arr[r] > pivot {
			r--
		}
		if l >= r {
			break
		}

		arr[l], arr[r] = arr[r], arr[l]

		if arr[l] == pivot {
			r--
		}
		if arr[r] == pivot {
			l++
		}
	}

	//在移动一下
	if l == r {
		l++
		r--
	}
	//向左递归
	if left < r {
		QuickSort(left, r, arr)
	}
	//向右递归
	if right > l {
		QuickSort(l, right, arr)
	}
}

//快速排序
func main() {
	arr := [6]int{-9, 78, 0, 23, -576, 70}
	fmt.Println(arr)
	QuickSort(0, len(arr)-1, &arr)
	fmt.Println("main")
	fmt.Println(arr)
}
