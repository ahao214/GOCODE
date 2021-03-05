package main

//题目：232.用栈实现队列
//说明：请你仅使用两个栈实现先入先出队列。队列应当支持一般队列的支持的所有操作（push、pop、peek、empty）
//输入：["MyQueue", "push", "push", "peek", "pop", "empty"]
//[[], [1], [2], [], [], []]
//输出： [null, null, null, 1, 1, false]
//说明：MyQueue myQueue = new MyQueue();
//myQueue.push(1); // queue is: [1]
//myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
//myQueue.peek(); // return 1
//myQueue.pop(); // return 1, queue is [2]
//myQueue.empty(); // return false
type MyQueue struct {
	Stack *[]int
	Queue *[]int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	tmpOne, tmpTwo := []int{}, []int{}
	return MyQueue{Stack: &tmpOne, Queue: &tmpTwo}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	*this.Stack = append(*this.Stack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(*this.Queue) == 0 {
		this.fromStackToQueue(this.Stack, this.Queue)
	}

	popped := (*this.Queue)[len(*this.Queue)-1]
	*this.Queue = (*this.Queue)[:len(*this.Queue)-1]
	return popped
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(*this.Queue) == 0 {
		this.fromStackToQueue(this.Stack, this.Queue)
	}

	return (*this.Queue)[len(*this.Queue)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(*this.Stack)+len(*this.Queue) == 0
}

func (this *MyQueue) fromStackToQueue(s, q *[]int) {
	for len(*s) > 0 {
		popped := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		*q = append(*q, popped)
	}
}
