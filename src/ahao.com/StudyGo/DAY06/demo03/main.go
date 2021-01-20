package main

import (
	"fmt"
)

//多个 defer 注册，按 FILO 次序执行 ( 先进后出 )。哪怕函数或某个延迟调用发生错误，这些调用依旧会被执行
func test(x int) {
	defer fmt.Println("a")
	defer fmt.Println("b")
	defer func() {
		fmt.Println(100 / x)
	}()
	defer fmt.Println("c")

}

//多defer注册
func main() {
	test(20)
}
