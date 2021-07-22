package demo03

import (
	"fmt"
)

//求和为n的所有整数组合
func getAllCombination(sum int, result []int, count int) {
	if sum < 0 {
		return
	}

	//数字的组合满足和为sum的条件，打印出所有组合
	if sum == 0 {
		fmt.Print("满足条件的组合：")
		for i := 0; i < count; i++ {
			fmt.Print(result[i], "")
			fmt.Println()
			return
		}
	}

	//打印debug信息，为了便于理解
	fmt.Print("---当前组合：---")
	for i := 0; i < count; i++ {
		fmt.Print(result[i], " ")
	}
	fmt.Println("------")

	//确定组合中下一个取值
	var next int
	if count == 0 {
		next = 1
	} else {
		next = result[count-1]
	}
	fmt.Println("---", "i=", next, "count=", count, "----") //打印debug信息，为了便于理解
	for i := next; i <= sum; i++ {
		result[count] = i
		count++
		getAllCombination(sum-i, result, count) //求和为sum-i的组合
		count--                                 //递归完成后，去掉最后一个组合的数字
	}
}

//方法功能：找出和为n的所有整数的组合
func shoWAllCombination(n int) {
	if n < 1 {
		fmt.Println("参数不满足要求")
		return
	}
	result := make([]int, n) //存储和为n的组合方式
	getAllCombination(n, result, 0)
}

//求正整数n所有可能的整数组合
func main() {
	shoWAllCombination(4)
}
