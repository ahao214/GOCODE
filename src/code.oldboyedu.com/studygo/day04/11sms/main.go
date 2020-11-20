package main

import(
	"fmt"
	"os"
)

/*
函数版学生管理系统
写一个系统可以查看、新增、删除学生
*/

type student struct{
	id int64
	name string
}

var (
	allStudent map[int64]*student	//变量声明

)

//构造函数
func newStudent(id int64,name string)*student{
	return &student{
		id:id,
		name:name,
	}
}

func showAllStudent(){
	for k,v:=range allStudent{
		fmt.Printf("学号：%d 姓名：%s\n",k,v.name)
	}
}

func addStudent(){
	//创建一个新学生
	var(
		id int64
		name string
	)
	fmt.Print("请输入学生的学号：")	
	fmt.Scanln(&id)
	fmt.Print("请输入学生的姓名：")
	fmt.Scanln(&name)
	//调用构造函数
	newStu:=newStudent(id,name)

	allStudent[id]=newStu
}

func deleteStudent(){
	var(
		deleteID int64
	)
	fmt.Print("请输入要删除的学生的学号")
	fmt.Scanln(&deleteID)
	//删除学生
	delete(allStudent,deleteID)



}


func main(){
	
	//初始化
	allStudent = make(map[int64]*student,50)
	for{
		fmt.Println("欢迎光临学生管理系统")
		fmt.Println(`
		1、查看所有学生
		2、新增学生
		3、删除学生
		4、退出
	`)
	fmt.Println("输入你想要干什么：")
	var choice int
	fmt.Scanln(&choice)
	fmt.Printf("你输入了%d这个选项\n",choice)

		switch(choice){
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1)	//退出
		default:
			fmt.Println("滚~")
		}
	}

}