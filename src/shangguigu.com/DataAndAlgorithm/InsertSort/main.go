package main

import (
	"fmt"
)

func InsertSort(arr *[5]int) {
	//完成第一次，给第二个元素找到合适的位置，并插入
	for j := 1; j < len(arr); j++ {
		insertVal := arr[j]
		insertIndex := j - 1
		//从大到小
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex] //数据后移
			insertIndex--
		}
		//插入
		if insertIndex+1 != j {
			arr[insertIndex+1] = insertVal
		}
		fmt.Printf("第%d次插入后：%v\n", j, *arr)
	}
}

//插入排序
func main() {
	arr := [5]int{23, 0, 12, 56, 34}
	fmt.Println("原始数组：", arr)
	InsertSort(&arr)
	fmt.Println("主函数：", arr)
}
