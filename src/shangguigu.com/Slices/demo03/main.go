package main

import(
	"fmt"
)

func main(){
	var arr[5]int=[...]int{10,20,30,40,50}
	//slice:=arr[1:4]	// 20 30 40
	slice:=arr[:]	// 10 20 30 40 50
	fmt.Println("for遍历\n")
	for i:=0;i<len(slice);i++{
		fmt.Printf("slice[%v]=%v\n",i,slice[i])
	}

	fmt.Println("for-range遍历\n")
	for i,v:=range slice{
		fmt.Printf("slice[%v]=%v\n",i,v)
	}

	//切片可以继续切片
	slice2:=slice[1:3]
	fmt.Println("slice2:",slice2)

	//切片动态追加
	var slice3[]int=[]int{100,200,300}

	slice3=append(slice3,400,500)
	fmt.Println("slice3:",slice3)

	//切片的拷贝
	//注意：1、拷贝必须都是切片 2、拷贝的两个切片相互独立，修改某个切片的值，不会影响另一个切片的值

	var a[]int=[]int{1,2,3,4,5}
	var copySlice=make([]int,10)
	fmt.Println("拷贝之前：",copySlice)
	fmt.Println(a)
	copy(copySlice,a)
	fmt.Println("拷贝之后：",copySlice)
	fmt.Println(a)
}