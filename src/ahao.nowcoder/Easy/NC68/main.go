package main

/**
 * 跳台阶
 * @param number int整型
 * @return int整型
 */
func jumpFloor(number int) int {
	if number == 0 {
		return 0
	}
	if number == 1 {
		return 1
	}
	if number == 2 {
		return 2
	}
	return jumpFloor(number-1) + jumpFloor(number-2)
}

func jumpFloorT(number int) int {
	if number <= 0 {
		return 0
	}
	if number == 1 {
		return 1
	}
	if number == 2 {
		return 2
	}
	arr := make([]int, number+1)
	arr[0] = 0
	arr[1] = 1
	arr[2] = 2
	for i := 3; i <= number; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[number]
}
