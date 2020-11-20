package main

import "fmt"

//类型断言
func assign(a interface{}){
	fmt.Printf("%T\n",a)
	str,ok:=a.(string)
	if !ok{
		fmt.Println("猜错了")
	}else{
		fmt.Println("传递进来的是一个字符串",str)
	}
	fmt.Println(str)
}

func assign2(a interface{}){
	fmt.Printf("%T\n",a)
	switch t:=a.(type){
	case string:
		fmt.Println("是一个字符串:",t)
	case int:
		fmt.Println("是一个int:",t)
	case bool:
		fmt.Println("是一个bool值：", t)
	}

}

func main(){
	assign(100)
	assign2(false)
	assign2("啊哈哈")
	assign2(int64(100))
}