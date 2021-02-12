package main

import (
	"fmt"
	"net"
	"os"
)

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误：", err.Error())
		os.Exit(1)
	}
}

//UDP客户端
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "用法：%s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkErr(err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkErr(err)
	_, err = conn.Write([]byte("anything"))
	checkErr(err)
	var buf [512]byte
	n, err := conn.Read(buf[0:])
	checkErr(err)
	fmt.Println(string(buf[0:n]))
	os.Exit(0)
}
