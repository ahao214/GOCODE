package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func echo(conn *net.TCPConn) {
	tick := time.Tick(5 * time.Second) //5秒请求一次
	for now := range tick {
		n, err := conn.Write([]byte(now.String()))
		if err != nil {
			log.Println(err)
			conn.Close()
			return
		}
		fmt.Printf("send %d bytes to %s\n", n, conn.RemoteAddr())
	}
}

func main() {
	address := net.TCPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8000,
	}
	listener, err := net.ListenTCP("tcp4", &address) //创建TCP4服务器端监听器
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Fatal(err) //错误直接退出
		}
		fmt.Println("远程地址：", conn.RemoteAddr())
		go echo(conn)
	}

}
