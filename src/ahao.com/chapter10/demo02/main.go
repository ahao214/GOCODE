package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

//通过关键字go启动一个goroutine
func main() {
	go say("world") //开一个新的goroutine执行
	say("hello")    //当前goroutine执行

}
