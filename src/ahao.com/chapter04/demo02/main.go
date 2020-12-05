package main

import(
	"fmt"
)


func main(){
	grade:="E"
	marks:=90

	switch{
	case marks>=90:
		grade="A"
	case marks>=80:
		grade="B"
	case marks>=70:
		grade="C"
	case marks>=60:
		grade="D"
	default:
		grade="E"
	}

	switch{
	case grade=="A":
		fmt.Printf("成绩优秀!\n")
	case grade=="B":
		fmt.Printf("表现良好!\n")
	case grade=="C", grade=="D":
		fmt.Printf("再接再厉！\n")
	/*fallthrough的写法*/
	case grade=="C":
		fallthrough
	case grade=="D":
		fmt.Printf("再接再厉！\n")

	default:
		fmt.Printf("成绩不合格！\n")
	}

	fmt.Printf("你的成绩是:%s\n",grade)
}