package main

import (
	"fmt"
)

/*
要求统计1~8000之间的素数
*/

//向intChan放入1~8000个数字
func putNum(intChan chan int) {
	for i := 1; i <= 8000; i++ {
		intChan <- i
	}
	//关闭
	close(intChan)
}

//开启4个协程，从intChan中取出数据，并判断是否是素数，如果是，就放入到primeChan中
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		num, ok := <-intChan
		if !ok { //intChan取不到...
			break
		}
		flag = true //假设是素数
		//判断num是否是素数
		for i := 2; i < num; i++ {
			if num%i == 0 { //说明该num不是素数
				flag = false
				break
			}
		}
		if flag {
			//将这个数放入到primeChan中
			primeChan <- num
		}
	}
	fmt.Println("有一个primeNum协程因为取不到数据，退出")
	//这里我们还不能关闭primeNum
	//向 exitChan写入true
	exitChan <- true
}

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000) //放入结果的管道
	//标识退出的管道
	exitChan := make(chan bool, 4)

	//开启1个协程，向intChan中放入1~8000的数字
	go putNum(intChan)

	//开启4个协程，从intChan中取出数据，并判断是否是素数，如果是，就放入到primeChan中
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	//这里我们主线程，要进行处理
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		//当我们从exitChan中取出4个结果，就可以放心的关闭primeChan
		close(primeChan)
	}()

	//遍历primeChan,取出结果
	for {
		v, ok := <-primeChan
		if !ok {
			break
		}
		//输出结果
		fmt.Printf("取出的素数是：%d\n", v)
	}
	fmt.Println("main线程退出")
}
