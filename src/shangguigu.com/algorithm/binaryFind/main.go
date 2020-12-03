package main

import(
	"fmt"
)

//二分查找
func BinaryFind(arr *[10]int,leftIndex int,rightIndex int, findVal int){
	
	if leftIndex>rightIndex{
		fmt.Println("找不到")
		return 
	}
	//先找到中间的下标
	middle:=(leftIndex+rightIndex)/2
	if(*arr)[middle]>findVal{
		//要找的数字在leftIndex和middle-1之间
		BinaryFind(arr,leftIndex,middle-1,findVal)
	}else if((*arr)[middle]<findVal){
		//要找的数字在middle+1和rightIndex之间
		BinaryFind(arr,middle+1,rightIndex,findVal)
	}else{
		fmt.Printf("找到了，下标为%v\n",middle)
	}

}



func main(){
	arr:=[10]int{2,4,11,23,46,67,70,78,89,100}
	BinaryFind(&arr,0,len(arr)-1,-46)	
}