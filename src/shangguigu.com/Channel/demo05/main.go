package main

import "fmt"

func main() {
	//声明管道为只写
	var writeChan chan<- int
	writeChan = make(chan int, 3)
	writeChan <- 100

	fmt.Println(writeChan)

	//声明管道为只读
	//var readChan <-chan int
	//num := <-readChan
	//fmt.Println(num)
}
