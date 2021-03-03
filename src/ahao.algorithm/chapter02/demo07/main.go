package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

//用于定义对象
type User struct {
	id   int
	name string
	seq  int
}

type LoginQueue struct {
	queue *gotype.SliceQueue
}

//进入队列尾部
func (p *LoginQueue) EnQueue(u *User) {
	p.queue.EnQueue(u)
	u.seq = p.queue.Size()
}

//出队列
func (p *LoginQueue) DeQueue() {
	p.queue.DeQueue()
	p.UpdateSeq()
}

//队列中的人随机离开
func (p *LoginQueue) DeQueueUser(u *User) {
	p.queue.Remove(u)
	p.UpdateSeq()
}

func (p *LoginQueue) UpdateSeq() {
	i := 1
	for _, v := range p.queue.List() {
		u := v.(*User)
		u.seq = i
		i++
	}
}

func (p *LoginQueue) PrintQueue() {
	for _, v := range p.queue.List() {
		fmt.Println(v.(*User))
	}
}

//如何设计一个排序系统
func main() {
	u1 := &User{id: 1, name: "user1"}
	u2 := &User{id: 1, name: "user2"}
	u3 := &User{id: 1, name: "user3"}
	u4 := &User{id: 1, name: "user4"}
	loginQueue := &LoginQueue{queue: gotype.NewSliceQueue()}
	loginQueue.EnQueue(u1)
	loginQueue.EnQueue(u2)
	loginQueue.EnQueue(u3)
	loginQueue.EnQueue(u4)

	loginQueue.DeQueue()       //队首元素u1出队列
	loginQueue.DeQueueUser(u3) //队列中间的元素u3出队列
	loginQueue.PrintQueue()
}
