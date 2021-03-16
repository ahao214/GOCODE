package main

import (
	"fmt"
)

func main() {
	d := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	one := d[3:7]
	two := one[1:3]
	for i := range two {
		two[i] += 10
	}
	fmt.Println(d)
	fmt.Println(one)
	fmt.Println(two)

}
