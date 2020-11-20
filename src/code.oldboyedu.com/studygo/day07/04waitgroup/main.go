package main

import(
	"fmt"
	"math/rand"
	"time"
	"sync"
)

//waitgroup

func f(){
	for i:=0;i<5;i++{
		rand.Seed(time.Now().UnixNano())
		r1:=rand.Int()		//int64
		r2:=rand.Intn(10) //10表示 0<=x<10
		fmt.Println(r1,r2)
	}
}

func f1(i int){
	defer wg.Done()
	time.Sleep(time.Second*time.Duration(rand.Intn(3)))
	fmt.Println(i)
}

var wg sync.WaitGroup

func main(){
	// f()

	for i:=0;i<10;i++{
		wg.Add(1)
		go f1(i)
	}
	wg.Wait()	//等待
}