package main

import (
	"fmt"
)

//循环语句range
func main() {
	s := "abc"
	for i := range s {
		fmt.Println(s[i])
	}

	for _, c := range s {
		fmt.Println(c)
	}

	m := map[string]int{"a": 1, "b": 2}
	//返回(key,value)
	for k, v := range m {
		fmt.Println(k, v)
	}

	//range会复制对象
	a := [3]int{0, 1, 2}
	for i, v := range a {
		if i == 0 { //修改前，先修改原数组
			a[1], a[2] = 888, 999
			fmt.Println(a)
		}
		a[i] = v + 100 //使用复制品中取出的value修改原数组
	}
	fmt.Println(a)

	one := []int{1, 2, 3, 4, 5}
	for i, v := range one {
		if i == 0 {
			one = one[:3]
			one[2] = 100 //对底层数据的修改
		}
		fmt.Println(i, v)
	}

}
