package main

import(
	"fmt"
)

//将斐波那契数列放到切片里面
func fbn(n int)([]uint64){
	//声明一个切片
	fbnSlice:=make([]uint64,n)
	fbnSlice[0]=1
	fbnSlice[1]=1
	for i:=2;i<n;i++{
		fbnSlice[i]=fbnSlice[i-1]+fbnSlice[i-2]
	}
	return fbnSlice
}

func main(){
	fbnSlice:=fbn(10)
	fmt.Println("fbnSlice:",fbnSlice)

}