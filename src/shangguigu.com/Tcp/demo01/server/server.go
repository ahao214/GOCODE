package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	//这里我们循环的接收客户端发送的数据
	defer conn.Close() //关闭conn

	for {
		//创建一个新的切片
		buf := make([]byte, 1024)
		//conn.Read(buf)
		//1.等待客户端通过conn发送信息
		//2.如果客户端没有发送，那么协程就阻塞在这里
		//fmt.Println("服务器在等待客户端的输入信息："+conn.RemoteAddr().String())
		n, err := conn.Read(buf) //从conn读取
		if err != nil {
			fmt.Println("客户端退出,err=", err)
			return
		}
		//3.显示客户端发送的内容到服务器的终端
		fmt.Print(string(buf[:n]))
	}
}

func main() {
	fmt.Println("服务器开始监听了...")
	//1.tcp表示使用的网络协议是tcp的
	//2. 0.0.0.0:8888 表示在本地监听8888端口
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}

	defer listen.Close() //延时关闭
	//循环等待客户端链接我
	for {
		//等待客户端链接
		fmt.Println("等待客户端来链接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
		} else {
			fmt.Printf("Accept() success=%v 客户端的ip是：%v\n", conn, conn.RemoteAddr().String())
		}
		//这里准备启动一个协程，为客户端服务
		go process(conn)
	}

	//fmt.Printf("listen success=%v\n",listen)

}
