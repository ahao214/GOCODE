package main

import "fmt"

//类型断言

func main(){
	var a interface{}
	a=100
	v1,ok:=a.(int8)
	if !ok{
		fmt.Println("猜对了，a是int8\n",v1)
		return 
	}else{
		fmt.Println("猜错了\n")
	}

	//方法2
	switch v2:=a.(type){
		case int8:
			fmt.Println("int8",v2)
		case int16:
			fmt.Println("int16",v2)
		case int64:
			fmt.Println("int64",v2)
		case int:
			fmt.Println("int",v2)
		default:
			fmt.Println("走人")
	}
	

}