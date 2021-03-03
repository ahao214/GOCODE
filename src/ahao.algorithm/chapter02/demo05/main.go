package main

import (
	"fmt"
	"github.com/isdamir/gotype"
	"math"
)

type MinStack struct {
	//用来存储栈中元素
	elemStack *gotype.SliceStack
	//栈顶永远存储当前elemStack中最小的值
	minStack *gotype.SliceStack
}

func (p *MinStack) Push(data int) {
	p.elemStack.Push(data)
	//更新保存最小元素的栈
	if p.minStack.IsEmpty() {
		p.minStack.Push(data)
	} else {
		if data <= p.minStack.Top().(int) {
			p.minStack.Push(data)
		}
	}
}

func (p *MinStack) Pop() int {
	topData := p.elemStack.Pop().(int)
	if topData == p.Min() {
		p.minStack.Pop()
	}
	return topData
}

func (p *MinStack) Min() int {
	if p.elemStack.IsEmpty() {
		return math.MaxInt32
	} else {
		return p.minStack.Top().(int)
	}
}

//如何用O(1)的时间复杂度求栈中最小元素
func main() {
	stack := &MinStack{elemStack: &gotype.SliceStack{}, minStack: &gotype.SliceStack{}}
	stack.Push(5)
	fmt.Println("栈中最小的值为：", stack.Min())
	stack.Push(6)
	fmt.Println("栈中最小的值为：", stack.Min())
	stack.Push(2)
	fmt.Println("栈中最小的值为：", stack.Min())
	stack.Pop()
	fmt.Println("栈中最小的值为：", stack.Min())
}
