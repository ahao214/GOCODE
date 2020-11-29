package main

import(
	"fmt"
	"time"
	"context"
	"go.etcd.io/etcd/clientv3"
)


//etcd watch 操作

func main(){
	cli,err:=New(clientv3.Config{
		Endpoints:[]string{"127.0.0.1:2379"},
		DialTimeout:5*time.Second,
	})

	//watch操作
	if err!=nil{
		fmt.Printf("connect etcd failed,err:%v\n",err)
		return
	}
	fmt.Println("connect etcd successed")

	defer cli.Close()

	//watch
	//安排一个哨兵一直监视着Tom这个key的变化(新增、删除、修改)
	ch:=cli.Watch(context.Background(),"Tom")
	
	//从通道尝试取值
	for wresp:=range ch{
		for _,evt:=range wresp.Events{
			fmt.Printf("Type:%v key:%v value:%v\n",evt.Type,evt.Kv.Key,evt.Kv.Value)

		}
	}

	
}