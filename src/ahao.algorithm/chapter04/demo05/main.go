package main

import "fmt"

//找出数组中出现奇数次的数
func getNum(arr []int) {
	if arr == nil || len(arr) < 1 {
		return
	}
	data := map[int]int{}
	for _, v := range arr {
		if vv, ok := data[v]; ok {
			if vv == 1 {
				data[v] = 0
			} else {
				data[v] = 1
			}
		} else {
			//map中没有这个数字，说明第一次出现，value赋值为1
			data[v] = 1
		}
	}

	for _, v := range arr { //使用数组来遍历，防止出现map的key随机分布
		if data[v] == 1 {
			fmt.Println(v)
		}
	}
}

func main() {
	fmt.Println("Hash 法")
	arr := []int{3, 5, 6, 6, 5, 7, 2, 2}
	getNum(arr)

}
