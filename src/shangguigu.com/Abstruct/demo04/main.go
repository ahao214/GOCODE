package main

import(
	"fmt"
)

type Student struct{
	Name string
	Age int
	Score int
}

func(stu *Student)ShowInfo(){
	fmt.Printf("学生姓名=%v 年龄=%v 成绩=%v",stu.Name,stu.Age,stu.Score)
}

func(stu *Student)SetScore(score int){
	stu.Score=score
}

type Pupil struct{
	Student		//嵌入了Student匿名结构体
}

func(p *Pupil)testing(){
	fmt.Println("小学生正在考试...")
}

type Graduate struct{
	Student		//嵌入了Student匿名结构体
}
func(g *Graduate)testing(){
	fmt.Println("大学生正在考试...")
}

func main(){

	p:=&Pupil{}
	p.Student.Name="jack"
	p.Student.Age=8
	p.testing()
	p.Student.SetScore(80)
	p.Student.ShowInfo()
}