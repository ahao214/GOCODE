package main

import "fmt"

func main() {
	array := []int{6, 8, 10}
	var res int
	res = min(array)
	fmt.Printf("最小值是：%d\n", res)

	fmt.Println("----")
	a, b := computer(2, 4)
	fmt.Println(a, b)
	fmt.Println("----")
	c := 1
	d := 2
	c, d = change(c, d)
	println(c, d)
	fmt.Println("----")
}

func change(a, b int) (x, y int) {
	x = a + 10
	y = b + 10
	return
}

func min(arr []int) (min int) {
	min = arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return
}

func computer(x, y int) (int, int) {
	return x + y, x * y
}
