package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

type StackQueue struct {
	aStack *gotype.SliceStack
	bStack *gotype.SliceStack
}

func (p *StackQueue) Push(data int) {
	p.aStack.Push(data)
}

func (p *StackQueue) Pop() int {
	if p.bStack.IsEmpty() {
		for !p.aStack.IsEmpty() {
			p.bStack.Push(p.aStack.Pop())
		}
	}
	return p.bStack.Pop().(int)
}

//用两个栈模拟队列操作
func main() {
	stack := &StackQueue{aStack: &gotype.SliceStack{}, bStack: &gotype.SliceStack{}}
	stack.Push(1)
	stack.Push(2)
	fmt.Println("队列首元素为：", stack.Pop())
	fmt.Println("队列首元素为：", stack.Pop())
}
