package main

import(
	"fmt"
)

//枚举

const 人数,费用,班级 = 40,200.02,"214班"
const(
	Monday,Tuesday,Wednesday=1,2,3
	Thursday,Friday,Saturday=4,5,6
)

const(
	a=iota
	b
	c
	d,e,f=iota,iota,iota
	g=iota
	h="h"
	i
	j=iota
)

const z=iota

func main(){
	
	fmt.Println(a,b,c,d,e,f,g,h,i,j,z)

}