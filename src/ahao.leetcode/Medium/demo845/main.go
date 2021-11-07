package main

//845. 数组中的最长山脉
func longestMountain(arr []int) int {
	left, right, res, isAscending := 0, 0, 0, true
	for left < len(arr) {
		if right+1 < len(arr) && ((isAscending == true && arr[right+1] > arr[left] && arr[right+1] > arr[right]) || (right != left && arr[right+1] < arr[right])) {
			if arr[right+1] < arr[right] {
				isAscending = false
			}
			right++
		} else {
			if right != left && isAscending == false {
				res = max(res, right-left+1)
			}
			left++
			if right < left {
				right = left
			}
			if right == left {
				isAscending = true
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestMountain2(arr []int) int {
	if arr == nil || len(arr) < 3 {
		return 0
	}
	slow, fast, res := 0, 0, 0
	for slow < len(arr) {
		if fast < len(arr)-1 && arr[fast+1] > arr[fast] {
			for fast < len(arr)-1 && arr[fast+1] > arr[fast] {
				fast++
			}
			if fast < len(arr)-1 && arr[fast+1] < arr[fast] {
				for fast < len(arr)-1 && arr[fast+1] < arr[fast] {
					fast++
				}
				if fast != slow && (fast-slow+1) > res {
					res = fast - slow + 1
				}
			}
		}
		if fast > slow {
			slow = fast
		} else {
			slow++
			fast = slow
		}
	}
	return res
}
