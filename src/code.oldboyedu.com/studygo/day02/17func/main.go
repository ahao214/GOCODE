package main


import "fmt"

//函数的定义,第一个括号写的是传入的参数，第二个括号写的是返回值
func sum(x int,y int)(ret int){
	return x + y
}

//没有返回值
func fone(x int,y int){
	fmt.Println(x + y)
}

//没有参数和返回值
func ftwo(){
	fmt.Println("没有参数和返回值")
}

//没有参数有返回值
func fthree() int{
	return 3
}

//多个返回值
func ffive()(int,string){
	return 1,"Tom"
}


func fsix(x,y int,n,m string, i,j bool) int{
	return 3
}



func f7(x int, y ...int){
	fmt.Println(x)
	fmt.Println(y)
}

func main(){
	r:= sum(1,2)
	fmt.Println(r)

	age,name :=ffive()
	fmt.Println(age)
	fmt.Println(name)

}