package main

import (
	"fmt"
	"math/rand"
	"time"
)

var arrone = [5]int{1, 2, 3, 4, 5}
var arrtwo = [...]int{1, 2, 3, 4, 5, 6}
var str = [5]string{3: "hello world", 4: "tom"}

//求元素和
func sumArr(a [10]int) int {
	var sum int = 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum
}

//找出数组中和为给定值的两个元素的下标，例如数组[1,3,5,8,7]，
//找出两个元素之和等于8的下标分别是（0，4）和（1，2）
//求元素和，是给定的值
func myTest(a [5]int, target int) {
	//遍历数组
	for i := 0; i < len(a); i++ {
		other := target - a[i]
		//继续遍历
		for j := i + 1; i < len(a); j++ {
			if a[j] == other {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
}

func main() {
	a := [3]int{1, 2}
	b := [...]int{1, 2, 3, 4}
	c := [5]int{2: 100, 4: 200}
	d := [...]struct {
		name string
		age  uint8
	}{
		{"userone", 10},
		{"usertwo", 20},
	}

	fmt.Println(arrone, arrtwo, str)
	fmt.Println(a, b, c, d)

	e := [...][2]int{{1, 1}, {2, 2}, {3, 3}} //第2维度不能使用"..."
	fmt.Println(e)

	//多维数组遍历
	var f [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}
	for k1, v1 := range f {
		for k2, v2 := range v1 {
			fmt.Printf("(%d,%d)=%d", k1, k2, v2)
		}
		fmt.Println()
	}

	// 若想做一个真正的随机数，要种子
	//seed()种子默认是1
	//rand.Seed(1)

	rand.Seed(time.Now().Unix())
	var nums [10]int
	for i := 0; i < len(nums); i++ {
		nums[i] = rand.Intn(1000)
	}
	sum := sumArr(nums)
	fmt.Printf("sum=%d\n", sum)

	arrNums := [5]int{1, 3, 5, 8, 7}
	myTest(arrNums, 8)

}
