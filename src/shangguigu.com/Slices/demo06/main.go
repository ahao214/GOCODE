package main

import(
	"fmt"
)

func main(){
	//string底层是个byte数组，因此string可以进行切片处理
	str:="hello@guigu"
	//使用切片获取guigu
	slice:=str[6:]
	fmt.Println("slice:",slice)

	//修改string里面的内容
	arr:=[]byte(str)
	arr[0]='z'
	str=string(arr)
	fmt.Println("str:",str)

	//将string转换为[]rune，就可以使用中文来修改
	arr1:=[]rune(str)
	arr1[0]='中'
	str=string(arr1)
	fmt.Println("str:",str)

}