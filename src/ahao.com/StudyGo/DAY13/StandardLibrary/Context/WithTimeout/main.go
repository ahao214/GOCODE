package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(ctx context.Context) {
LOOP:
	for {
		fmt.Println("db connecting...")
		time.Sleep(time.Millisecond * 10)
		select {
		case <-ctx.Done(): //50毫秒后自动调用
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
	wg.Done()
}

// WithTimeout
//WithTimeout返回WithDeadline(parent, time.Now().Add(timeout))。
//取消此上下文将释放与其相关的资源，
// 因此代码应该在此上下文中运行的操作完成后立即调用cancel，
// 通常用于数据库或者网络连接的超时控制
func main() {
	//设置一个50毫秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel() //通知子goroutine结束
	wg.Wait()
	fmt.Println("over")

}
