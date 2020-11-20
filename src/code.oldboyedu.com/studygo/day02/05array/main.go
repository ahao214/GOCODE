package main

import "fmt"

func main(){
	//数组
	var a1[3] bool
	var a2[4] bool

	fmt.Println("a1:%T a2:%T\n", a1, a2)

	//数组的初始化
	a1 = [3]bool{true,true,true}
	fmt.Println(a1)

	//根据初始值自动推断数组的长度
	a10:=[...]int{0,1,2,3,4,5,6,7,8}
	fmt.Println(a10)
	fmt.Println(len(a10))

	//根据索引来初始化
	a3:=[5]int{0:1,4:2}
	fmt.Println(a3)


	//数组的变量
	citys := [...]string{"北京","上海","广州","深圳"}

	for i:=0;i<len(citys);i++{
		fmt.Println(citys[i])
	}

	//for range 变量数组
	for i,v:=range citys{
		fmt.Println(i,v)
	}

	//多维数组
	var a11[3][2]int
	a11 =[3][2]int{
		[2]int{1,2},
		[2]int{3,4},
		[2]int{5,6},
	}

	for _,v1:= range a11{
		fmt.Println(v1)
		for _,v2:=range v1{
			fmt.Println(v2)
		}
	}


	s1 := [...]int{1,3,5,7,8}
	sum :=0
	for _,v:=range s1{
		sum+=v
	}
	fmt.Println(sum)


	for i:=0;i<len(s1);i++{
		for j:=i+1;j<len(s1);j++{
			if s1[i]+s1[j]==8{
				fmt.Printf("(%d,%d)\n",i,j)
			}
		}
	}

}