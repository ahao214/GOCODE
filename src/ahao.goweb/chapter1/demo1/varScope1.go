package main

import (
	"fmt"
)

func main() {
	//声明局部变量
	var localone, localtwo, localthree int
	//初始化参数
	localone = 2
	localtwo = 4
	localthree = localone * localtwo

	fmt.Printf("localone =%d,localtwo=%d,localthree=%d\n", localone, localtwo, localthree)

}
