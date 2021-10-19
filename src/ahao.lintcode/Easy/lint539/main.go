package main

/**
 * 539 · 移动零
 * @param nums: an integer array
 * @return: nothing
 */
func moveZeroes(nums *[]int) {
	if len(*nums) == 0 {
		return
	}
	j := 0
	for i := 0; i < len(*nums); i++ {
		if (*nums)[i] != 0 {
			if i != j {
				(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
				j++
			} else {
				j++
			}
		}
	}
}
