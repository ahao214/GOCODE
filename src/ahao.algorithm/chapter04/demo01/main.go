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

//异或方法
func FindDupByXOR(arr []int) int {
	if arr == nil {
		return -1
	}
	result := 0
	len := len(arr)
	for _, v := range arr {
		result ^= v
	}
	for i := 1; i < len; i++ {
		result ^= i
	}
	return result
}

//数据映射方法
func FindDupByMap(arr []int) int {
	if arr == nil {
		return -1
	}
	len := len(arr)
	index := 0
	i := 0
	for true {
		//数组中的元素的值只能小于len，否则会溢出
		if arr[i] >= len {
			return -1
		}
		if arr[index] < 0 {
			break
		}
		//访问过，通过变相反数的方法进行标记
		arr[index] *= -1
		//index的后续为array[index]
		index = arr[index] * -1
		if index >= len {
			fmt.Println("数组中有非法数字")
			return -1
		}
	}
	return index

}

//环形相遇方法
func FindDupByLoop(arr []int) int {
	if arr == nil {
		return -1
	}
	slow := 0
	fast := 0
	for ok := true; ok; ok = slow != fast {
		fast = arr[arr[fast]] //fast一次走两步
		slow = arr[slow]      //slow一次走一步
	}
	fast = 0
	//找到入口点
	for ok := true; ok; ok = slow != fast {
		fast = arr[fast]
		slow = arr[slow]
	}
	return slow
}

func main() {
	arr := []int{1, 3, 4, 2, 5, 3}
	fmt.Println("Hash方法")
	fmt.Println(FindDupByHash(arr))

	fmt.Println("异或方法")
	fmt.Println(FindDupByXOR(arr))

	fmt.Println("数据映射方法")
	fmt.Println(FindDupByMap(arr))

	fmt.Println("环形相遇方法")
	fmt.Print(FindDupByLoop(arr))
}
