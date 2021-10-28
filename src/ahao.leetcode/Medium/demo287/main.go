package main

import "fmt"

//287. 寻找重复数
func findDuplicate(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)>>1
		count := 0
		for _, num := range nums {
			if num <= mid {
				count++
			}
		}
		if count > mid {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func findDuplicateT(nums []int) int {
	slow, fast := nums[0], nums[0]
	slow = nums[slow]
	fast = nums[nums[fast]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	p1 := nums[0]
	p2 := slow
	for p1 != p2 {
		p1 = nums[p1]
		p2 = nums[p2]
	}
	return p1
}

func main() {
	nums := []int{1, 3, 4, 2, 2}
	fmt.Println(findDuplicate(nums))
}
