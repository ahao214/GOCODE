package main

/**
 * NC140 排序
 * 将给定数组排序
 * @param arr int整型一维数组 待排序的数组
 * @return int整型一维数组
 */
func MySort(arr []int) []int {
	llen := len(arr)
	for i := 0; i < llen-1; i++ {
		for j := llen - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				tmp := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = tmp
			}
		}
	}
	return arr
}

func MySortQ(arr []int) []int {
	return sort(arr, 0, len(arr)-1)
}

func sort(arr []int, low, high int) []int {
	if low >= high {
		return arr
	}

	i := low
	j := high
	var index int
	index = arr[i]
	for i < j {
		for i < j && arr[j] >= index {
			j--
		}
		if i < j {
			arr[i] = arr[j]
			i++
		}
		for i < j && arr[i] < index {
			i++
		}
		if i < j {
			arr[j] = arr[i]
			j--
		}
	}
	arr[i] = index
	arr = sort(arr, low, i-1)
	arr = sort(arr, i+1, high)
	return arr
}
