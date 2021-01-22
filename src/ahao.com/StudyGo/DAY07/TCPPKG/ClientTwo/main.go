package main

import (
	proto "ahao.com/StudyGo/DAY07/TCPPKG/ProtoTwo"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed,err:", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `hello,hello.how are you?`
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed,err:", err)
			return
		}
		conn.Write(data)
	}

}
