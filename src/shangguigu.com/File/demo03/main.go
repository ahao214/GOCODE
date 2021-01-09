package main

import(
	"fmt"
	"io/ioutil"
)

func main(){
	//使用ioutil.ReadFile一次性将文件读取出来
	file:="f:/one.txt"
	content,err:=ioutil.ReadFile(file)
	if err!=nil{
		fmt.Printf("read file err=%v\n",err)
	}
	//在终端将内容输出
	fmt.Printf("%v",content)	//[]byte
	fmt.Printf("%v",string(content))

}