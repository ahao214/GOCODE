package main

import (
	"fmt"
)

//按要求构造新的数组
func calculate(a, b []int) {
	b[0] = 1
	ll := len(a)
	for i := 1; i < ll; i++ {
		b[i] = b[i-1] * a[i-1]
	}
	b[0] = a[ll-1]
	for i := ll - 2; i >= 1; i-- {
		b[i] *= b[0]
		b[0] *= a[i]
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := make([]int, len(a))
	calculate(a, b)
	fmt.Println(b)

}
