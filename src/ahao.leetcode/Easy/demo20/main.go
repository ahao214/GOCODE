package main

import "fmt"

//题目：20.有效的括号
//说明：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//输入：s = "()"
//输出：true
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	stack := make([]rune, 0)
	for _, v := range s {
		if (v == '[') || (v == '(') || (v == '{') {
			stack = append(stack, v)
		} else if ((v == ']') && len(stack) > 0 && stack[len(stack)-1] == '[') ||
			((v == ')') && len(stack) > 0 && stack[len(stack)-1] == '(') ||
			((v == '}') && len(stack) > 0 && stack[len(stack)-1] == '{') {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0

}

func isValid2(s string) bool {
	if len(s) == 0 {
		return true
	}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if isLeft(s[i]) {
			stack = append(stack, s[i])
		} else {
			if len(stack) > 0 && stack[len(stack)-1] == findLeft(s[i]) {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func isLeft(c byte) bool {
	if c == '(' || c == '[' || c == '{' {
		return true
	}
	return false
}

func findLeft(c byte) byte {
	if c == ')' {
		return '('
	}
	if c == ']' {
		return '['
	}
	if c == '}' {
		return '{'
	}
	return ' '
}

func main() {
	s := "(){}}}}"
	fmt.Println(isValid(s))
}
