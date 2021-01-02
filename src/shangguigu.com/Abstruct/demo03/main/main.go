package main

import(
	"fmt"
	"shangguigu.com/Abstruct/demo03/model"
)


func main(){
	account:=model.NewAccount("62220200","123456",40.4)
	if account!=nil{
		fmt.Println("创建成功...",account)
	}else{
		fmt.Println("创建失败...")
	}

}