package main

import (
	"bytes"
	"fmt"
)

//消除字符串的内嵌括号

//去掉字符串中嵌套的括号
func removeNestedPare(str string) string {
	arr := []rune(str)
	Parentheses_num := 0 //用来记录不匹配的"("出现的次数
	if arr[0] != '(' || arr[len(arr)-1] != ')' {
		return ""
	}
	var buf bytes.Buffer
	buf.WriteRune('(')
	//字符串首尾的括号可以单独处理
	for _, v := range arr {
		switch v {
		case '(':
			Parentheses_num++
		case ')':
			Parentheses_num--
		default:
			buf.WriteRune(v)
		}
	}

	//判断括号是否匹配
	if Parentheses_num != 0 {
		fmt.Println("由于括号不匹配，因此不做任何操作")
		return ""
	}
	//处理字符串结尾的')'
	buf.WriteRune(')')
	return buf.String()
}

func main() {
	str := "(1,(2,3),(4,(5,6),7))"
	fmt.Printf("%v去除嵌套括号后为：%v", str, removeNestedPare(str))

}
