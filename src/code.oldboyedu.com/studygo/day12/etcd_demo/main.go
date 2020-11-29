package main

import(
	"fmt"
	"time"
	"context"
	"go.etcd.io/etcd/clientv3"
)

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

	//put
	ctx,cancel:=context.WithTimeout(context.Background(),time.Second)
	_,err=cli.Put(ctx,"qimi","dsb")
	cancel()
	if err!=nil{
		fmt.Printf("put to etcd failed,err:%v\n",err)
		return
	}

	//get
	ctx,cancel=context.WithTimeout(context.Background(),time.Second)
	resp,err:=cli.Get(ctx,"qimi",Clientv3.WithPrefix())
	cancel()
	if err!=nil{
		fmt.Printf("get to etcd failed,err:%v\n",err)
		return 
	}

	for _,ev:=range resp.Kvs{
		fmt.Printf("%s:%s\n",ev.Key,ev.Value)
	}
}