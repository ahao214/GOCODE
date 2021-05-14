package main

//查询数组中元素的最大值和最小值

import (
	"fmt"
)

//分冶方法
func GetMaxAndMin(arr []int) (max, min int) {
	if arr == nil {
		return 0, 0
	}
	len := len(arr)
	max = arr[0]
	min = arr[0]
	//两两分组，把较小的数放到左半部分，较大的放到有半部分
	for i := 0; i < len-1; i = i + 2 {
		if arr[i] > arr[i+1] {
			tmp := arr[i]
			arr[i] = arr[i+1]
			arr[i+1] = tmp
		}
	}

	//在各个分组的左半部分找最小值
	for i := 2; i < len; i = i + 2 {
		if arr[i] < min {
			min = arr[i]
		}
	}
	//在各个分组的有半部分找最大值
	for i := 3; i < len; i = i + 2 {
		if arr[i] > max {
			max = arr[i]
		}
	}

	//如果数组中元素个数是个奇数个，最后一个元素被分为一组，需要特殊处理
	if len%2 == 1 {
		if max < arr[len-1] {
			max = arr[len-1]
		}
		if min > arr[len-1] {
			min = arr[len-1]
		}
	}

	return
}

//变形的分冶方法
func GetMaxAndMinRe(arr []int, l, r int) (max, min int) {
	if arr == nil {
		return
	}
	//求中点
	m := (l + r) / 2
	//l与r之间只有一个元素
	if l == r {
		max = arr[l]
		min = arr[l]
		return
	}
	//l与r之间只有两个元素
	if l+1 == r {
		if arr[l] >= arr[r] {
			max = arr[l]
			min = arr[r]
		} else {
			max = arr[r]
			min = arr[l]
		}
		return
	}

	//递归计算左半部分
	lmax, lmin := GetMaxAndMinRe(arr, l, m)

	//递归计算有半部分
	rmax, rmin := GetMaxAndMinRe(arr, m+1, r)
	//总的最大值
	if lmax > rmax {
		max = lmax
	} else {
		max = rmax
	}
	if lmin < rmin {
		min = lmin
	} else {
		min = rmin
	}

	return
}

func main() {
	arr := []int{7, 3, 19, 40, 4, 7, 1}
	max, min := GetMaxAndMin(arr)
	fmt.Println("分冶法")
	fmt.Println("最大值:", max)
	fmt.Println("最小值:", min)

	fmt.Println("扩展分冶方法")
	rmax, rmin := GetMaxAndMinRe(arr, 0, len(arr)-1)
	fmt.Println("最大值:", rmax)
	fmt.Println("最小值:", rmin)

}
