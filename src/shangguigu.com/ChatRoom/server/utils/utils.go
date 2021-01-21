package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"shangguigu.com/ChatRoom/common/message"
)

//这里将这些方法关联到结构体中
type Transfer struct {
	//分析它应该有哪些字段
	Conn net.Conn
	Buf  [8096]byte //这是传输时，使用缓冲
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	//buf := make([]byte, 8094)
	fmt.Println("等待客户端发送的数据...")
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		//err=errors.New("read pkg header error")
		return
	}

	//根据我读到的buf[:4]长度，转成一个uint32
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])

	//根据pkgLen读取消息内容
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//err=errors.New("read pkg body error")
		return
	}

	//把pkgLen反序列化成->message.Message
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("反序列化失败,err=", err)
		return
	}

	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	//var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	//发送长度
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write buf err=", err)
		return
	}

	// 发送data本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write data err=", err)
		return
	}
	return
}
