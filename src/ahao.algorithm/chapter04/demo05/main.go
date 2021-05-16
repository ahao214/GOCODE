package main

import "fmt"

//找出数组中出现奇数次的数-Hash 法
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

//异或法
func getNumXor(arr []int) {
	if arr == nil || len(arr) < 1 {
		return
	}
	result := 0
	position := uint(0)
	//计算数组中所有数字异或的结果
	for _, v := range arr {
		result ^= v
	}
	tmpResult := result
	//找出异或结果中其中一个位值为1的位数(如1100，位值为1位数为2和3)
	for i := result; i&1 == 0; i = i >> 1 {
		position++
	}
	//异或的结果与所有第position位为1的数异或，结果一定是出现一次的两个数中其中一个
	for _, v := range arr {
		if (v>>position)&1 == 1 {
			result ^= v
		}
	}
	fmt.Println(result)
	//得到另外一个出现一次的数
	fmt.Println(result ^ tmpResult)
}

func main() {
	fmt.Println("Hash 法")
	arr := []int{3, 5, 6, 6, 5, 7, 2, 2}
	getNum(arr)
	fmt.Println("异或法")
	getNumXor(arr)
}
