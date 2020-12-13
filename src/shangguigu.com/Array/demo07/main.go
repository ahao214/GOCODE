package main

import(
	"fmt"
)

func main(){

	var myChars[26] byte
	for i:=0;i<26;i++{
		myChars[i]='A'+byte(i)	//需要将i转成byte类型
	}

	for i:=0;i<26;i++{
		fmt.Printf("%c ", myChars[i])
	}


	intArr:=[...]int{2,45,23,12,100,89,46,67}
	maxVal:=intArr[0]
	maxValIndex:=0
	for i:=1;i<len(intArr);i++{
		if maxVal<intArr[i]{
			maxVal=intArr[i]
			maxValIndex=i
		}
	}
	fmt.Printf("最大值是:%v 最大值的下标是:%v\n",maxVal,maxValIndex)

}