package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandomM(a []int, n, m int) {
	if a == nil || n <= 0 || n < m {
		fmt.Println("参数不合法")
		return
	}
	for i := 0; i < m; i++ {
		j := i + rand.Intn(n-i) //获取i到n-1间的随机数
		//随机选出的元素放到数组的前面
		tmp := a[i]
		a[i] = a[j]
		a[j] = tmp
	}
}

//如何等概率地从大小为n的数组中选取m个整数
func main() {
	rand.Seed(time.Now().UnixNano())
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n, m := 10, 6
	getRandomM(a, n, m)
	for i := 0; i < m; i++ {
		fmt.Print(a[i], " ")
	}
	fmt.Println()

}
