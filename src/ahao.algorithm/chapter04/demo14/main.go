package main

import "fmt"

//求集合的所有子集

//位图法
func getAllSybset(arr []rune, mask []int, c int) {
	len := len(arr)
	if len == c {
		fmt.Print("{")
		for i, v := range arr {
			if mask[i] == 1 {
				fmt.Print(string(v), " ")
			}
		}
		fmt.Println("}")
	} else {
		mask[c] = 1
		getAllSybset(arr, mask, c+1)
		mask[c] = 0
		getAllSybset(arr, mask, c+1)
	}
}

//迭代法
func getAllSybsetSub(str string) (arr []string) {
	arr = make([]string, 0)
	if str == "" {
		return arr
	}
	arr = append(arr, string(str[0]))
	for i := 1; i < len(str); i++ {
		ll := len(arr)
		for j := 0; j < ll; j++ {
			arr = append(arr, arr[j]+string(str[i]))
		}
		arr = append(arr, string(str[i:i+1]))
	}
	return arr
}

func main() {
	arr := []rune{'a', 'b', 'c'}
	mask := []int{0, 0, 0}
	fmt.Println("位图法")
	getAllSybset(arr, mask, 0)

	fmt.Println("迭代法")
	result := getAllSybsetSub("abc")
	for _, v := range result {
		fmt.Println(v)
	}
}
