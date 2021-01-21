package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"shangguigu.com/ChatRoom/common/message"
)

//写一个函数，完成登录校验
func login(userId int, userPwd string) (err error) {
	//开始定协议

	//fmt.Printf(" userId = %d userPwd = %s\n", userId, userPwd)
	//
	//return nil

	//1.链接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}

	//延时关闭
	defer conn.Close()

	//2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType

	//3.创建一个LoginMes 结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//4.将LoginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//5.把data赋给了mes.Data字段
	mes.Data = string(data)

	//6.将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//7.到这个时候，data就是我们要发送的信息
	//7.1 先把data的长度发送给服务器
	//先获取到data的长度->转出一个表示长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}

	fmt.Println("客户端,发送消息的长度OK")

	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) err=", err)
		return
	}

	//这里还需要处理服务器端返回的消息
	mes, err = readPkg(conn) //mes 就是
	if err != nil {
		fmt.Println("readPkg er=", err)
		return
	}
	//将mes.Data 反序列化成LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)

	if loginResMes.Code == 200 {
		fmt.Println("登录成功")
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	return
}
