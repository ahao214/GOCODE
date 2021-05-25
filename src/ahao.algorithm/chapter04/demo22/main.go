package main

//从三个有序数组中找出它们的公共元素
import (
	"fmt"
)

func FindCommon(arrone, arrtwo, arrthree []int) {
	none := len(arrone)
	ntwo := len(arrtwo)
	nthree := len(arrthree)
	//遍历三个数组
	for i, j, k := 0, 0, 0; i < none && j < ntwo && k < nthree; {
		//找到了公共元素
		if arrone[i] == arrtwo[j] && arrtwo[j] == arrthree[k] {
			fmt.Print(arrone[i], " ")
			i++
			j++
			k++
		} else if arrone[i] < arrtwo[j] {
			i++
		} else if arrtwo[j] < arrthree[k] {
			j++
		} else {
			k++
		}
	}
}

func main() {
	arrone := []int{2, 5, 12, 20, 45, 85}
	arrtwo := []int{16, 19, 20, 85, 200}
	arrthree := []int{3, 4, 15, 20, 39, 72, 85, 190}
	FindCommon(arrone, arrtwo, arrthree)
}
