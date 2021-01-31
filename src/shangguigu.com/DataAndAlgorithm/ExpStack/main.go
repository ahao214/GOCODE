package main

import (
	"errors"
	"fmt"
	"strconv"
)

//使用数组模拟一个栈的使用
type Stack struct {
	MaxTop int     //表示栈最大可以存放数个数
	Top    int     //表示栈顶
	arr    [20]int //数组模拟栈
}

//入栈
func (this *Stack) Push(val int) (err error) {
	//先判断栈是否已经满了
	if this.Top == this.MaxTop-1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}
	this.Top++
	//放入数据
	this.arr[this.Top] = val
	return
}

//出栈
func (this *Stack) Pop() (val int, err error) {
	if this.Top == -1 {
		fmt.Println("stack empty")
		return -1, errors.New("stack empty")
	}
	val = this.arr[this.Top]
	this.Top--
	return val, nil
}

//遍历栈 需要从栈顶开始
func (this *Stack) Show() {
	//先判断栈是否为空
	if this.Top == -1 {
		fmt.Println("stack empty")
		return
	}
	fmt.Println("栈的情况如下：")
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, this.arr[i])
	}
}

//判断一个字符是不是运算符[+,-,*,/]
func (this *Stack) IsOper(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}

//计算
func (this *Stack) Cal(num1 int, num2 int, oper int) int {
	result := 0
	switch oper {
	case 42:
		result = num2 * num1
	case 43:
		result = num2 + num1
	case 45:
		result = num2 - num1
	case 47:
		result = num2 / num1
	default:
		fmt.Println("运算符错误")
	}
	return result
}

//返回某个运算符的优先级[程序员定的]
func (this *Stack) Priority(oper int) int {
	res := 0
	if oper == 42 || oper == 47 {
		res = 1
	} else if oper == 43 || oper == 45 {
		res = 0
	}
	return res
}

func main() {
	//数字栈
	numStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}
	//符号栈
	operStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}
	exp := "30+2*6-2"
	//定义一个index，帮助扫描exp
	index := 0

	//为了配合运算，定义需要的变量
	num1 := 0
	num2 := 0
	oper := 0
	res := 0
	keepNum := ""

	for {
		//这里需要增加一个逻辑，处理多位数

		ch := exp[index : index+1]
		temp := int([]byte(ch)[0])  //字符对应的ASCII码
		if operStack.IsOper(temp) { //说明是符号
			//如果operSatck是空栈，直接入栈
			if operStack.Top == -1 { //空栈
				operStack.Push(temp)
			} else {
				//如果operStack栈顶的运算符的优先级大于等于当前准备入栈的运算符的优先级
				//就从符号栈pop出，并从数字栈pop出两个数字，进行运算，运算后的结果在重新入栈到数字栈,
				//符号在入符号栈
				if operStack.Priority(operStack.arr[operStack.Top]) >= operStack.Priority(temp) {
					num1, _ = numStack.Pop()
					num2, _ = numStack.Pop()
					oper, _ = operStack.Pop()
					res = operStack.Cal(num1, num2, oper)
					//将计算的结果重新入数字栈
					numStack.Push(res)
					//当前的符号压入到符号栈
					operStack.Push(temp)
				} else {
					//当前的符号压入到符号栈
					operStack.Push(temp)
				}
			}
		} else { //说明是数字
			//处理多位数字
			//1.定义一个变量 keepNum 做拼接
			keepNum += ch
			//2.每次要向index的后面字符测试一下，看看是不是运算符 然后处理
			//如果已经到表达式的最后，直接将keepNum
			if index == len(exp)-1 {
				val, _ := strconv.ParseInt(keepNum, 10, 64)
				numStack.Push(int(val))
			} else {
				//向index后面测试看看是不是运算符
				if operStack.IsOper(int([]byte(exp[index+1 : index+2])[0])) {
					val, _ := strconv.ParseInt(keepNum, 10, 64)
					numStack.Push(int(val))
					keepNum = ""
				}
			}
			//val, _ := strconv.ParseInt(ch, 10, 64)
			//numStack.Push(int(val))
		}
		//继续扫描
		//先判断index是否已经扫描到计算表达式的最后
		if index+1 == len(exp) {
			break
		}
		index++
	}
	//如果扫描表达式完毕，依次从符号栈取出符号，从数字栈取出两个数字
	//运算后的结果入数字栈，直到符号栈为空
	for {
		if operStack.Top == -1 { //退出条件
			break
		}
		num1, _ = numStack.Pop()
		num2, _ = numStack.Pop()
		oper, _ = operStack.Pop()
		res = operStack.Cal(num1, num2, oper)
		//将计算的结果重新入数字栈
		numStack.Push(res)
	}
	//结果是numStack最后的数字
	res, _ = numStack.Pop()
	fmt.Printf("表达式%s=%v", exp, res)
}
