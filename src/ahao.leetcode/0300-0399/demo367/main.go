package main

//367. 有效的完全平方数
func isPerfectSquare(num int) bool {
	low, high := 1, num
	for low <= high {
		mid := low + (high-low)>>1
		if mid*mid == num {
			return true
		} else if mid*mid < num {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}

func isPerfectSquare2(num int) bool {
	x0 := float64(num)
	for {
		x1 := (x0 + float64(num)/x0) / 2
		if x0-x1 < 1e-6 {
			x := int(x0)
			return x*x == num
		}
		x0 = x1
	}
}
