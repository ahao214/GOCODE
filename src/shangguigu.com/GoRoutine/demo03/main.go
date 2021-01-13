package main

import (
	"fmt"
	"sync"
	"time"
)

//需要计算1-200各个数字的阶乘，并把结果放到一个map里面
//最后显示出来，要使用goroutine来完成

//思路
//1.编写一个函数，来计算各个数的阶乘，并放入到一个map中
//2.我们的启动的协程多个，统计的将结果放入到map中
//3.map 应该做一个全局的

//声明一个全局的map
var (
	myMap = make(map[int]int, 10)
	//声明一个全局的互斥锁
	lock sync.Mutex
)

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	//将res放入到myMap中
	//加锁
	lock.Lock()
	myMap[n] = res
	//解锁
	lock.Unlock()
}

func main() {

	//开启多个协程
	for i := 1; i <= 200; i++ {
		go test(i)
	}

	//休眠10秒钟
	time.Sleep(time.Second * 10)
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("myMap[%d]=%d\n", i, v)
	}
	lock.Unlock()

}
