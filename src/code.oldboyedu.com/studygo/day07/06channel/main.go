package main

import(
	"fmt"
	"sync"
)

var a[] int
var b chan int	//需要指定通道中元素的类型
var wg sync.WaitGroup

func noBufChannel(){
	b=make(chan int)	//不带缓冲区通道的初始化
	wg.Add(1)
	go func(){
		defer wg.Done()
		x:=<-b
		fmt.Println("后台goroutine从通道b中取到了:",x)
	}()
	b<-10
	fmt.Println("10发送到了通道b中...")
	wg.Wait()

}

func BufChannel(){
	b=make(chan int,1)	//带缓冲区的通道的初始化
	b<-20
	fmt.Println("20发送到了通道b中...")
	x:=<-b
	fmt.Println("从通道b中取到了:",x)
	close(b)	//关闭通道b
}

func main(){
	BufChannel()


}