package main

//1300 · 巴什博弈

/**
 * @param n: an integer
 * @return: whether you can win the game given the number of stones in the heap
 */
func canWinBash(n int) bool {
	// Write your code here
	if n%4 == 0 {
		return false
	}
	return true
}
