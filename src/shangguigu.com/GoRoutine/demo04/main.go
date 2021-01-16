package main

import (
	"fmt"
)

func main() {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan) //close
	//这时不能在写入channel
	//intChan<-300
	//当管道关闭后，读数据是可以的
	n := <-intChan
	fmt.Println("n:", n)

	//channel的遍历
	intChantwo := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChantwo <- i * 2
	}

	//遍历管道，不能使用普通的for循环结构
	//使用for-range遍历

	//在遍历的时候，如果channel没有关闭，则会出现deadlock的错误

	close(intChantwo)
	for v := range intChantwo {
		fmt.Println("v=", v)
	}

	fmt.Println(intChan)

}
