package main

import "fmt"


type student struct{
	id int64
	name string
}

//造一个学生管理者
type studentMgr struct{
	allStudent map[int64]student
}

//查看学生
func(s studentMgr)showStudents(){
	for _,stu:=range s.allStudent{
		fmt.Printf("学号:%d,姓名:%s\n",stu.id,stu.name)
	}
}

//增加学生
func(s studentMgr) addStudent(){
	var(
		stuID int64
		stuName string
	)

	fmt.Print("请输入学号：")
	fmt.Scanln(&stuID)
	fmt.Print("请输入姓名：")
	fmt.Scanln(&stuName)
	//根据用户输入创建结构体对象
	newStu:=student{
		id:stuID,
		name:stuName,
	}
	//把新的学生放到s.allStudent这个map中
	s.allStudent[newStu.id]=newStu
	fmt.Println("添加成功！")
}

//修改学生
func(s studentMgr) editStudent(){
	//1、获取学生输入的学号
	var stuID int64
	fmt.Print("请输入学号：")
	fmt.Scanln(&stuID)

	//2、展示该学号对应的学生信息，如果没有则提示不存在
	stuObj,ok:=s.allStudent[stuID]
	if !ok{
		fmt.Println("查无此人")
		return
	}
	fmt.Printf("你需要修改的学生信息如下，学号：%d,姓名：%s\n",stuObj.id,stuObj.name)

	//3、请输入修改后的学生名
	fmt.Println("请输入学生的新名字：")
	var newName string
	fmt.Scanln(&newName)

	//4、更新学生的姓名
	stuObj.name = newName
	s.allStudent[stuID]=stuObj	//更新map中的学生
}

//删除学生
func (s studentMgr) deleteStudent(){
	//1.输入要删除的学生的学号
	var stuID int64
	fmt.Print("请输入要删除的学生的学号：")
	fmt.Scanln(&stuID)
	//2.查找学号是否存在，不存在，提示查无此人
	_,ok:=s.allStudent[stuID]
	if !ok{
		fmt.Println("查无此人")
		return
	}

	//3.如果有的话，就从map中删除键值对
	delete(s.allStudent,stuID)
	fmt.Println("删除成功！")
}