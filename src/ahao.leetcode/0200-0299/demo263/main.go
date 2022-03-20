package main

import "fmt"

//263. 丑数
func isUgly(n int) bool {
	if n > 0 {
		for _, i := range []int{2, 3, 5} {
			for n%i == 0 {
				n /= i
			}
		}
	}
	return n == 1
}

func main() {
	fmt.Println(isUgly(6))
}
