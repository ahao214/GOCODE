package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"shangguigu.com/ChatRoom/common/message"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8094)
	fmt.Println("等待客户端发送的数据...")
	_, err = conn.Read(buf[:4])
	if err != nil {
		return
	}

	//根据我读到的buf[:4]长度，转成一个uint32
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])

	//根据pkgLen读取消息内容
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		return
	}

	//把pkgLen反序列化成->message.Message
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("反序列化失败,err=", err)
		return
	}

	return
}

func writePkg(conn net.Conn, data []byte) (err error) {
	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write buf err=", err)
		return
	}

	// 发送data本身
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write data err=", err)
		return
	}
	return
}
