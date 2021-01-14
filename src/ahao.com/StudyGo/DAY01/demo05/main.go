package main

import (
	"fmt"
	"math"
)

//遍历字符串
func traversalString() {
	s := "pprof.cn博客"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v(%c)", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s {
		fmt.Printf("%v(%c)", r, r)
	}
	fmt.Println()
}

//修改字符串
func changeString() {
	s := "hello"
	//强制类型转换
	byteS := []byte(s)
	byteS[0] = 'H'
	fmt.Println(string(byteS))

	s2 := "博客"
	runeS := []rune(s2)
	runeS[0] = '微'
	fmt.Println(string(runeS))
}

//类型转换
func sqrtDemo() {
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func main() {
	traversalString()
	changeString()
	sqrtDemo()
}
