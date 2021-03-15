package main

import "fmt"

func test(s string, a ...int) {
	fmt.Printf("%T,%v\n", a, a) //显示类型和值
}

func testone(a ...int) {
	for i := range a {
		a[i] += 10
	}
}

//变参
//变参本质就是一个切片，只能接收一到多个同类型参数，且必须放到列表尾部
func main() {
	test("abc", 1, 2, 3, 4)

	a := []int{10, 20, 30}
	testone(a...)
	fmt.Println(a)

}
