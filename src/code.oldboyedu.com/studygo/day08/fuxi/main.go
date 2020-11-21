package main


import(
	"fmt"
	"time"
	"math/rand"
)

//channel


//往通道中发送至
func sendNum(ch chan int){
	for{
		num:=rand.Intn(10)
		ch<-num
		time.Sleep(5*time.Second)
	}
}

func main(){
	ch:=make(chan int,1)
	//ch<-100	//把一个值发送到通道中
	//ch<-200
	// <-ch	//把通道中的100取出来
	// x,ok:=<-ch	//再次取值就会阻塞
	// fmt.Println(x,ok)
	go sendNum(ch)
	for{
		x,ok:=<-ch	//阻塞等4秒
		fmt.Println(x,ok)
		time.Sleep(time.Second)
	}

}