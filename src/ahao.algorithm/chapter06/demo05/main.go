package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

//产生的随机数是整数1·7的均匀分布
func rand7() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(7)
}

//产生的随机数是整数1·10的均匀分布
func rand10() int {
	x := math.MaxInt32
	for x > 40 {
		x = (rand7())*7 + rand7()
	}
	return x%10 + 1
}

//根据已知随机数生成函数计算新的随机数
func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(rand10())
	}
}
