package main

import (
	"fmt"
)

func main() {
	switch x := 5; x {
	default:
		fmt.Println(x)
	case 5:
		x += 10
		fmt.Println(x)
		fallthrough //继续执行下一个case，但不在匹配条件表达式
	case 6:
		x += 20
		fmt.Println(x)
	}
}
