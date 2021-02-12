package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func handleClient(conn *net.UDPConn) {
	var buf [512]byte
	_, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}
	daytime := time.Now().String()
	conn.WriteToUDP([]byte(daytime), addr)
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误：", err.Error())
		os.Exit(1)
	}
}

//UDP服务器端
func main() {
	service := ":1200"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkErr(err)
	conn, err := net.ListenUDP("udp", udpAddr)
	checkErr(err)
	for {
		handleClient(conn)
	}

}
