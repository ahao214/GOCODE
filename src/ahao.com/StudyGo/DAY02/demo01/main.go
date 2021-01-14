package main

import (
	"fmt"
)

//创建切片的各种方式

func main() {
	//1.声明切片
	var one []int
	if one == nil {
		fmt.Println("是空")
	} else {
		fmt.Println("不是空")
	}

	//2. :=
	two := []int{}
	//3.make()
	var three []int = make([]int, 0)
	fmt.Println(one, two, three)

	//4.初始化赋值
	var four []int = make([]int, 0, 0)
	fmt.Println(four)
	five := []int{1, 2, 3}
	fmt.Println(five)

	//5.从数组切片
	arr := [5]int{1, 2, 3, 4, 5}
	var six []int
	//前包后不包
	six = arr[1:4]
	fmt.Println(six)

}
