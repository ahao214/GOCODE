package main

import (
	"errors"
	"fmt"
	"os"
)

//使用一个结构体来管理环形队列
type CircleQueue struct {
	maxValue int    //4
	array    [5]int //数组
	head     int    //指向队列的队首
	tail     int    //指向队列的队尾
}

//入队列
func (this *CircleQueue) Push(val int) (err error) {
	if this.IsFull() {
		return errors.New("queue is full")
	}
	//分析this.tail 在队列尾部，但是不包含最后的元素
	this.array[this.tail] = val //把值给尾部
	this.tail = (this.tail + 1) % this.maxValue
	return
}

//出队列
func (this *CircleQueue) Pop() (val int, err error) {
	if this.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	//取出 head是指向队首，并且含队首元素
	val = this.array[this.head]
	this.head = (this.head + 1) % this.maxValue
	return
}

//显示队列
func (this *CircleQueue) ListQueue() {
	fmt.Println("环形队列情况如下：")
	//取出当前队列有多少个元素
	size := this.Size()
	if size == 0 {
		fmt.Println("队列为空")
	}
	//设计一个辅助变量，指向head
	tempHead := this.head
	for i := 0; i < size; i++ {
		fmt.Printf("array[%d]=%d\t", tempHead, this.array[tempHead])
		tempHead = (tempHead + 1) % this.maxValue
	}
	fmt.Println()
}

//判断环形队列是否满了
func (this *CircleQueue) IsFull() bool {
	return (this.tail+1)%this.maxValue == this.head
}

//判断环形列表是否为空
func (this *CircleQueue) IsEmpty() bool {
	return this.tail == this.head
}

//取出环形队列有多少个元素
func (this *CircleQueue) Size() int {
	//这是一个关键的算法
	return (this.tail + this.maxValue - this.head) % this.maxValue
}

//数组模拟环形队列
func main() {
	//初始化一个环形队列
	queue := &CircleQueue{
		maxValue: 5,
		head:     0,
		tail:     0,
	}
	var key string
	var val int
	for {
		fmt.Println("1.输入add表示添加数据到队列")
		fmt.Println("2.输入get表示从队列中获取数据")
		fmt.Println("3.输入show表示显示队列中的数据")
		fmt.Println("4.输入exit表示退出")
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队列的数字")
			fmt.Scanln(&val)
			err := queue.Push(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列ok")
			}
		case "get":
			val, err := queue.Pop()
			if err != nil {
				fmt.Println("err:", err)
			} else {
				fmt.Println("从队列中获取到的值是：", val)
			}
		case "show":
			queue.ListQueue()

		case "exit":
			os.Exit(0)
		default:
			fmt.Println("你输入错误，请重新输入")
		}
	}
}
