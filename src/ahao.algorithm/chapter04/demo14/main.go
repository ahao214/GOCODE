package main

import "fmt"

//求集合的所有子集

//位图法
func getAllSybset(arr []rune, mask []int, c int) {
	len := len(arr)
	if len == c {
		fmt.Print("{")
		for i, v := range arr {
			if mask[i] == 1 {
				fmt.Print(string(v), " ")
			}
		}
		fmt.Println("}")
	} else {
		mask[c] = 1
		getAllSybset(arr, mask, c+1)
		mask[c] = 0
		getAllSybset(arr, mask, c+1)
	}
}

func main() {
	arr := []rune{'a', 'b', 'c'}
	mask := []int{0, 0, 0}
	fmt.Println("位图法")
	getAllSybset(arr, mask, 0)

}
