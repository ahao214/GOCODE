package main

//231.2的幂
func isPowerOfTwo(n int) bool {
	return n&(n-1) == 0
}
