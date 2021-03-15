package main

import (
	"log"
)

//ConnLimiter 定义一个结构体
type ConnLimiter struct {
	concurrentConn int
	bucket         chan int
}

//NewConnLimiter ...
func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: cc,
		bucket:         make(chan int, cc),
	}
}

//GetConn 获取通道里面的值
func (cl *ConnLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.concurrentConn {
		log.Printf("Reached the rate limitation.")
		return false
	}

	cl.bucket <- 1
	return true
}

//ReleaseConn 释放通道里面的值
func (cl *ConnLimiter) ReleaseConn() {
	c := <-cl.bucket
	log.Printf("New connction coming: %d", c)
}
