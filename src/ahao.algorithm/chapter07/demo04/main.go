package main

import (
	"fmt"
	"math/rand"
	"time"
)

func func1() int {
	return rand.Intn(2)
}

//返回0的概率为1/4，返回1的概率为3/4
func func2() int {
	a := func1()
	b := func1()
	tmp := a
	tmp |= (b << 1)
	if tmp == 0 {
		return 0
	} else {
		return 1
	}
}

//如何用一个随机函数得到另外一个随机函数
func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 16; i++ {
		fmt.Print(func2())
	}
	fmt.Println()
	for i := 0; i < 16; i++ {
		fmt.Print(func2())
	}
	fmt.Println()
}
