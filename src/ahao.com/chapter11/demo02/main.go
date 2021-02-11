package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

func checkSum(msg []byte) uint16 {
	sum := 0
	//先假设为偶数
	for n := 1; n < len(msg)-1; n += 2 {
		sum += int(msg[n])*256 + int(msg[n+1])
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)
	var answer uint16 = uint16(^sum)
	return answer
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误：%s", err.Error())
		os.Exit(1)
	}
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()
	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil

}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("用法：", os.Args[0], "hostname")
		os.Exit(1)
	}

	service := os.Args[1]
	conn, err := net.Dial("ip4:icmp", service)
	checkError(err)
	var msg [512]byte
	msg[0] = 8
	msg[1] = 0
	msg[2] = 0
	msg[3] = 0
	msg[4] = 0
	msg[5] = 13
	msg[6] = 8
	msg[7] = 37
	len := 8
	check := checkSum(msg[0:len])
	msg[2] = byte(check >> 8)
	msg[3] = byte(check & 255)

	_, err = conn.Write(msg[0:len])
	checkError(err)

	_, err = conn.Read(msg[0:])
	checkError(err)

	fmt.Println("取得响应")
	if msg[5] == 13 {
		fmt.Println("Identifier匹配")
	}
	if msg[7] == 37 {
		fmt.Println("Sequence匹配")
	}
	os.Exit(0)

}
