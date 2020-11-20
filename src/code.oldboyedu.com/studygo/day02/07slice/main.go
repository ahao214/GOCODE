package main

import "fmt"

func main(){
	//切片

	//定义切片
	var s1[]int
	var s2[]string
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	//初始化
	s1 = []int{1,2,3}
	s2=[]string{"北京","上海","广州","深圳"}
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s1 ==nil)
	fmt.Println(s2==nil)

	//切片的长度和容量
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1),cap(s1))
	fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2),cap(s2))

	//从数组获取到切片
	a1:=[...]int{1,3,5,7,9,11,13}
	s3:=a1[0:4]	//左包含右不包含
	fmt.Println(s3)
	

}