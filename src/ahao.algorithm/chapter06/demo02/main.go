package main

import "fmt"

//判断n能否表示成2的n此方法
func isPower1(n int) bool {
	if n < 1 {
		return false
	}
	for i := 1; i < n; {
		if i == n {
			return true
		}
		i <<= 1
	}
	return false
}

func isPower2(n int) bool {
	if n < 1 {
		return false
	}
	return n&(n-1) == 0

}

//判断一个数是否为2的n次方
func main() {
	fmt.Println("方法一：")
	if isPower1(9) {
		fmt.Println("9能表示成2的n次方")
	} else {
		fmt.Println("9不能表示成2的n次方")
	}

	if isPower2(8) {
		fmt.Println("8能表示成2的n次方")
	} else {
		fmt.Println("8不能表示成2的n次方")
	}
}
