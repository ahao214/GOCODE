package main


import "fmt"

func f1(f func()){
	fmt.Println("this is f1")
	f()
}

func f2(x,y int){
	fmt.Println(x+y)
	fmt.Println("this is f2")
}


//定义一个函数对f2进行包装
func f3(x,y int) func(){
	tmp:=func(){
		fmt.Println(x+y)
	}
	return tmp
}

//f func(int,int)是第一个参数
//x和y分别是第二个和第三个参数
func f4(f func(int,int),x,y int) func(){
	tmp:=func(){
		f(x,y)
	}
	return tmp
}


func main(){
	ret:=f3(100,200)
	f1(ret)

	ret2:=f4(f2,200,300)
	f1(ret2)
	
}