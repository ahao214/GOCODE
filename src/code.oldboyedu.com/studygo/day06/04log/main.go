package main

import (
	"log"
	"time"
	"fmt"
	"os"

)

func main(){
	//往文件里面写日志
	fileObj,err:=os.OpenFile("./xx.log",os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err!=nil{
		fmt.Printf("open file failed,err:%v\n",err)
		return 
	}
	log.SetOutput(fileObj)

	for{
		log.Println("这是一条普通的日志")
		time.Sleep(time.Second*3)		
	}
}