package main

import (
	"errors"
	"fmt"
)

//使用数组模拟一个栈的使用
type Stack struct {
	MaxTop int    //表示栈最大可以存放数个数
	Top    int    //表示栈顶
	arr    [5]int //数组模拟栈
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

func main() {
	stack := &Stack{
		MaxTop: 5,  //表示最多存放5个数到栈中
		Top:    -1, //当栈顶为-1时，表示栈为空
		arr:    [5]int{},
	}
	//入栈
	for i := 1; i <= stack.MaxTop; i++ {
		stack.Push(i)
	}

	//出栈
	val, _ := stack.Pop()
	fmt.Println("出栈的值：", val)

	//显示
	stack.Show()

}
