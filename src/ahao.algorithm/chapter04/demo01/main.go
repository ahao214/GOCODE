package main

//找出数组中唯一的重复元素
import "fmt"

//Hash方法
func FindDupByHash(arr []int) int {
	if arr == nil {
		return -1
	}
	data := map[int]bool{}
	for _, v := range arr {
		if _, ok := data[v]; ok {
			return v
		} else {
			data[v] = true
		}
	}
	return -1
}

func main() {
	arr := []int{1, 3, 4, 2, 5, 3}
	fmt.Println("Hash方法")
	fmt.Println(FindDupByHash(arr))

}
