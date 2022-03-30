package main

//507. 完美数
func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}

	sum := 1
	for d := 2; d*d <= num; d++ {
		if num%d == 0 {
			sum += d
			if d*d < num {
				sum += num / d
			}
		}
	}
	return sum == num
}

func checkPerfectNumber2(num int) bool {
	return num == 6 || num == 28 || num == 496 || num == 8128 || num == 33550336
}
