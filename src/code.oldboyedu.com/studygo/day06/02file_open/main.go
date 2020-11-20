package main

import (
	"fmt"
	"os"

)


func f1(){
	var fileObj *os.File
	var err error

	fileObj,err=os.Open("./main.go")
	if err!=nil{
		fmt.Printf("open file failed,err:%v\n",err)
	}
	defer fileObj.Close()

}

func f2(){
	fileObj,err:=os.OpenFile("./sb.txt",os.O_RDWR,0644)
	if err!=nil{
		fmt.Printf("open file failed,err:%v\n",err)
		return 
	}
	fileObj.Close()

	fileObj.Seek(1,0)	//光标的位置
	s=[]byte{'c'}
	fileObj.Write(s)


	// var ret[1] byte
	// n,err:=fileObj.Read(ret[:])
	// if err!=nil{
	// 	fmt.Printf("read from file failed,err:%v\n",err)
	// 	return 
	// }
	// fmt.Println(string(ret[:n]))


}

func main(){

	f2()
}