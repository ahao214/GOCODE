package main

import (
	"fmt"
	"time"
)

type MyTime struct {
	time.Time //匿名字段
}

func (t MyTime) first3Chars() string {
	return t.Time.String()[0:3]
}

func (t MyTime) first4Chars() string {
	return t.Time.String()[0:4]
}

func main() {
	m := MyTime{time.Now()}
	//调用匿名Time上的string方法
	fmt.Println("完整的时间格式：", m.String())

	//调用MyTime.first3Chars
	fmt.Println("前三个字符：", m.first3Chars())

	fmt.Println("前四个字符：", m.first4Chars())
}
