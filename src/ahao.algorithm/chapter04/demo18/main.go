package main

import (
	"fmt"
	. "github.com/isdamir/gotype"
)

//判断请求是否在给定的存储条件下完成
//按照R[i]-O[i]由大到小进行操作
func bubbleSort(R, O []int) {
	ll := len(R)
	for i := 0; i < ll-1; i++ {
		for j := ll - 1; j > i; j-- {
			if R[j]-O[j] > R[j-1]-O[j-1] {
				SwapInt(R, j, j-1)
				SwapInt(O, j, j-1)
			}
		}
	}
}

func schedule(R, O []int, M int) bool {
	bubbleSort(R, O)
	//剩余可用的空间数
	left := M
	for i, v := range R {
		//剩余的空间无法继续处理dii个请求
		if left < v {
			return false
		} else {
			//剩余的空间能继续处理第i个请求，处理完成后将占用O[i]个空间
			left -= O[i]
		}
	}
	return true
}

func main() {
	R := []int{10, 15, 23, 20, 6, 9, 7, 16}
	O := []int{2, 7, 8, 4, 5, 8, 6, 8}
	scheduleResult := schedule(R, O, 50)
	if scheduleResult {
		fmt.Println("按照如下请求序列可以完成：")
		for i := 0; i < 8; i++ {
			fmt.Print(R[i], ",", O[i], " ")
		}
	} else {
		fmt.Print("无法完成调度")
	}
}
