package main

import(
	"fmt"
)

func modify(map1 map[int]int){
	map1[10]=900
}

//定义一个学生结构体
type Stu struct{
	Name string
	Age int
	Address string
}

func main(){
	//map是引用类型
	//修改后，会直接修改原来的map
	map1:=make(map[int]int)

	map1[1]=100
	map1[2]=200
	map1[10]=10
	map1[20]=20
	modify(map1)

	fmt.Println(map1)


	students:=make(map[string]Stu,10)
	stu1:=Stu{"Tom",23,"北京"}
	stu2:=Stu{"Merry",25,"上海"}
	students["no1"]=stu1
	students["no2"]=stu2
	fmt.Println(students)

	//遍历各个学生的信息
	for k,v:=range students{
		fmt.Printf("学生的编号是:%v \n",k)
		fmt.Printf("学生的姓名是:%v \n",v.Name)
		fmt.Printf("学生的年龄是:%v \n",v.Age)
		fmt.Printf("学生的地址是:%v \n",v.Address)
		fmt.Println()
	}
}