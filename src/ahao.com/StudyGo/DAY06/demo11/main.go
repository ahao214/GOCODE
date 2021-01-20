package main

import (
	"fmt"
)

type Person interface {
	Speak(string) string
}

type Student struct {
}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是大帅比"
	} else {
		talk = "你好"
	}
	return
}

func main() {
	//var per Person=Student{}	//报错
	var peo = Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
