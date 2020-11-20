package main


import(
	"fmt"
	"time"
	"math/rand"
	"sync"
)


type job struct{
	value int64
}

type result struct{
	job *job
	sum int64
}

var jobChan=make(chan *job,100)
var resultChan=make(chan *result,100)
var wg sync.WaitGroup

func zhoulin(zl chan<- *job){
	defer wg.Done()
	//循环生成int64位的随机数，发送到jobChan
	for{
		x:=rand.Int63()
		newJob:=&job{
			value:x,
		}
		zl<-newJob
		time.Sleep(time.Second)
	}
}


func baodelu(zl<-chan *job, resultChan chan <- *result){
	defer wg.Done()
	for{
		job:=<-zl
		sum:=int64(0)
		n:=job.value
		for n>0{
			sum+=n%10
			n=n/10
		}
		newResult:=&result{
			job:job,
			sum:sum,
		}
		resultChan<-newResult
	}
}

func main(){
	wg.Add(1)
	go zhoulin(jobChan)
	for i:=0;i<24;i++{
		go baodelu(jobChan,resultChan)
	}
	for result:=range resultChan{
		fmt.Printf("value:%d sum:%d\n",result.job.value,result.sum)
	}
	wg.Wait()
}