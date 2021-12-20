package main

import (
	"math"
	"sort"
)

//475.供暖器

//排序+二分
func findRadius(houses, heaters []int) (ans int) {
	sort.Ints(heaters)
	for _, house := range houses {
		j := sort.SearchInts(heaters, house+1)
		minDis := math.MaxInt32
		if j < len(heaters) {
			minDis = heaters[j] - house
		}
		i := j - 1
		if i >= 0 && house-heaters[i] < minDis {
			minDis = house - heaters[i]
		}
		if minDis > ans {
			ans = minDis
		}
	}
	return
}

//排序+双指针
func findRadius2(houses, heaters []int) (ans int) {
	sort.Ints(houses)
	sort.Ints(heaters)
	j := 0
	for _, house := range houses {
		dis := abs(house - heaters[j])
		for j+1 < len(heaters) && abs(house-heaters[j]) >= abs(house-heaters[j+1]) {
			j++
			if abs(house-heaters[j]) < dis {
				dis = abs(house - heaters[j])
			}
		}
		if dis > ans {
			ans = dis
		}
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
