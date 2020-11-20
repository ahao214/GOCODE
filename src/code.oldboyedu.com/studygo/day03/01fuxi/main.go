package main

import "fmt"

func main(){
	var ages[30] int
	fmt.Println(ages)
	ages = [30]int{1,2,34,56}
	var age2=[...]int{56,7,9,10,11,12}
	fmt.Println(age2)

	var age3 =[...]int{1:100,10:999}
	fmt.Println(age3)

	//多维数组只有最外层可以使用...
	var a1[3][2]int 
	a1=[3][2]int{
		[2]int{1,2},
		[2]int{3,4},
		[2]int{5,6},
	}
	fmt.Println(a1)

	x:=[3]int{1,2,3}
	fmt.Println(x)
	f1(x)
	fmt.Println(x)

}

func f1(a[3]int){
	a[1]=100
}