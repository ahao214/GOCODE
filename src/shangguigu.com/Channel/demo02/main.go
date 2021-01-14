package main

import (
	"fmt"
)

type Cat struct {
	Name string
	Age  int
}

func main() {
	//定义一个可以存在任何类型的管道
	var allChan chan interface{}
	allChan = make(chan interface{}, 3)
	allChan <- 10
	allChan <- "Tom And Jerry"
	cat := Cat{"amao", 12}
	allChan <- cat

	//我们希望获取到管道中的第三个元素，则先将前两个推出
	<-allChan
	<-allChan
	//从管道中取出cat
	newCat := <-allChan
	fmt.Printf("newCat=%T,newCat=%v\n", newCat, newCat)
	//使用类型断言
	a := newCat.(Cat)
	fmt.Printf("newCat.Name=%v\n", a.Name)
	fmt.Printf("newCat.Age=%v\n", a.Age)

}
