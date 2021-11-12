package main

//JZ11 旋转数组的最小数字
func minNumberInRotateArray(rotateArray []int) int {
	low := 0
	high := len(rotateArray) - 1
	for low < high {
		pivot := low + (high-low)/2
		if rotateArray[pivot] < rotateArray[high] {
			high = pivot
		} else if rotateArray[pivot] > rotateArray[high] {
			low = pivot + 1
		} else {
			high--
		}
	}
	return rotateArray[low]
}
