package main

import (
	"fmt"
)

type data int

func (d data) String() string {
	return fmt.Sprintf("data:%d", d)
}

//类型转换
func main() {
	var d data = 15
	var x interface{} = d
	if n, ok := x.(fmt.Stringer); ok { //转为更具体的接口类型
		fmt.Println(n)
	}
	if d2, ok := x.(data); ok { //转换回原始类型
		fmt.Println(d2)
	}
	e := x.(error)
	fmt.Println(e)
}
