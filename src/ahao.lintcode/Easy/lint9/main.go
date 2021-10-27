package main

import "fmt"

/**
 * 9 · Fizz Buzz 问题
 * @param n: An integer
 * @return: A list of strings.
 */
func fizzBuzz(n int) []string {
	res := make([]string, n)
	for i := 1; i <= n; i++ {
		switch {
		case (i%15 == 0):
			res[i-1] = "fizz buzz"
		case i%3 == 0:
			res[i-1] = "fizz"
		case i%5 == 0:
			res[i-1] = "buzz"
		default:
			res[i-1] = fmt.Sprintf("%d", i)
		}
	}
	return res
}
