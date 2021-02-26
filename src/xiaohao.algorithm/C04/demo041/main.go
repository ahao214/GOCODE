package main

import "fmt"

//Nim游戏
func canWinNim(n int) bool {
	return n%4 != 0
}

func main() {
	n := 5
	fmt.Println(canWinNim(n))
}
