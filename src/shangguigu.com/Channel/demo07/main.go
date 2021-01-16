package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello world")
	}
}

func test() {
	//这里可以使用defer+recover
	defer func() {
		//捕获test抛出的panic
		if err := recover(); err != nil {
			fmt.Println("test()协程发生错误", err)
		}
	}()
	var myMap map[int]string
	myMap[0] = "golang" //error

}

func main() {
	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("主函数在说话：", i)
		time.Sleep(time.Second)
	}

}
