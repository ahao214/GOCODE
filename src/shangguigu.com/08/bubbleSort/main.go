package main

import(
	"fmt"
)

//冒泡排序
func BubbleSort(arr *[5]int){
	fmt.Println("排序前的数组:",*arr)
	//临时变量
	temp:=0
	for i:=0;i<len(*arr)-1;i++{
		for j:=0;j<len(*arr)-1-i;j++{
			if (*arr)[j]>(*arr)[j+1]{
				temp=(*arr)[j]
				(*arr)[j]=(*arr)[j+1]
				(*arr)[j+1]=temp
			}
		}
	}
	
	fmt.Println("排序后的数组:",*arr)
	
}

func main(){
	//定义数组
	arr:=[5]int{12,45,1,56,6}
	BubbleSort(&arr)

	fmt.Println("main arr:",arr)	//有序

}