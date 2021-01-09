package main

import(
	"fmt"
)


//可变参数
func main(){
	//手动填写参数
	age:=ageMinOrMax("min",1,2,4,0,9)
	fmt.Printf("年龄最小的是%d岁\n",age)
	//数组作为参数
	ageArr:=[]int{9,1,12,89,6}
	age=ageMinOrMax("max",ageArr...)
	fmt.Printf("年龄最大的是%d岁\n",age)
}


func ageMinOrMax(s string,a ...int)int{
	if len(a)==0{
		return 0
	}

	if s=="max"{
		max:=a[0]
		for _,v:=range a{
			if v>max{
				max=v
			}
		}
		return max
	}else if(s=="min"){
		min:=a[0]
		for _,v:=range a{
			if v<min{
				min=v
			}
		}
		return min
	}else{
		e:=-1
		return e
	}
}