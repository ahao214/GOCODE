package main

import (
	"fmt"
)

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a)

	//append函数
	var runes []rune
	for _, r := range "hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)

}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
