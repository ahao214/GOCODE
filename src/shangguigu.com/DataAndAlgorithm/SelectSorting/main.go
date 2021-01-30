package main

import (
	"fmt"
)

//完成排序
func SelectSort(arr *[5]int) {

	for j := 0; j < len(arr)-1; j++ {
		//假设arr[0]就是最大值
		max := arr[j]
		maxIndex := j
		//遍历后面的1-[len(arr)-1]
		for i := j + 1; i < len(arr); i++ {
			if max < arr[i] { //找到真正的最大值
				max = arr[i]
				maxIndex = i
			}
		}
		//交换
		if maxIndex != j {
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
		}
		fmt.Printf("第%d次 %v \n", j+1, *arr)
	}
}

//选择排序
func main() {
	//定义一个数组
	arr := [5]int{10, 34, 19, 100, 80}
	fmt.Println("原来的数组：", arr)
	SelectSort(&arr)
}
