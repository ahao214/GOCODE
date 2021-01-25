package main

import (
	"errors"
	"fmt"
	"os"
)

//使用一个结构体管理队列
type Queue struct {
	maxSize int
	array   [5]int //数组=>模拟队列
	front   int    //表示指向队列首部
	rear    int    //表示指向队列尾部
}

//添加数据到队列
func (this *Queue) AddQueue(val int) (err error) {
	//先判断队列是否已满
	if this.rear == this.maxSize-1 { //重要的提示；rear是队列的尾部(含最后元素)
		return errors.New("queue full")
	}
	this.rear++ //rear 后移
	this.array[this.rear] = val

	return
}

//显示队列,找到队首，然后遍历到队尾
func (this *Queue) ShowQueue() {
	fmt.Println("队列的当前情况是：")
	//this.front 不包含队首的元素
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("array[%d]=%d\t", i, this.array[i])
	}
	fmt.Println()
}

//从队列中取数据
func (this *Queue) GetQueue() (val int, err error) {
	//先判断队列是否为空
	if this.rear == this.front { //队列为空
		return -1, errors.New("queue empty")
	}
	this.front++
	val = this.array[this.front]
	return val, err
}

//数组模拟队列实现
func main() {
	//先创建一个队列
	queue := Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
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
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列ok")
			}
		case "get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println("err:", err)
			} else {
				fmt.Println("从队列中获取到的值是：", val)
			}
		case "show":
			queue.ShowQueue()

		case "exit":
			os.Exit(0)
		default:
			fmt.Println("你输入错误，请重新输入")
		}
	}

}
