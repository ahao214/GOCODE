package main

import(
	"fmt"
	"os"
)

func writeDemo1(){
	fileObj,err:=os.OpenFile("./xx.txt",os.O_APPEND|os.O_CREATE,0644)
	if err!=nil{
		fmt.Printf("open file failed,err:%v",err)
		return 
	}

	//Write
	fileObj.Write([]byte("zhoulin shi yige hao xuesheng\n"))
	fileObj.WriteString("www.jd.com")
	fileObj.Close()
}

func main(){

	//writeDemo1

	

}