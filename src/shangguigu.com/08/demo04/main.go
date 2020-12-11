package main

import(
	"fmt"
	"strings"
)

func makeSuffix(suffix string)func(string)string{
	return func(name string)string{
		//如果name没有指定的后缀名，则加上，否则就返回原来的名字
		if !strings.HasSuffix(name,suffix){
			return name+suffix
		}
		return name
	}
}


func main(){
	//返回一个闭包
	f:=makeSuffix(".jpg")
	fmt.Println("文件名处理后=",f("winter"))
	fmt.Println("文件名处理后=",f("summer.jpg"))


}