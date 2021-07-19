package main

import (
	"fmt"
	"math/rand"
	"time"
)

//总共n个房间，判断用指定的策略是否能拿到最多金币
func getMaxNum(n int) bool {
	if n < 1 {
		fmt.Println("参数不合法")
		return false
	}
	a := make([]int, n)
	//随机生成n个房间里金币的个数
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(n) + 1 //生成1-n的随机数
	}
	//找出前四个房间中最多的金币个数
	max4 := 0
	for i := 0; i < 4; i++ {
		if a[i] > max4 {
			max4 = a[i]
		}
	}

	for i := 4; i < n-1; i++ {
		if a[i] > max4 {
			//能拿到最多的金币
			return true
		}
	}
	return false //不能拿到最多的金币

}

//如何拿到最多金币
func main() {
	rand.Seed(time.Now().UnixNano())
	monitorCount := 1000
	success := 0.0
	for i := 0; i < monitorCount; i++ {
	}
	if getMaxNum(10) {
		success++
	}
	fmt.Println(success / float64(monitorCount))

}
