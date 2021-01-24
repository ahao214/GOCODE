package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"shangguigu.com/ChatRoom/client/utils"
	"shangguigu.com/ChatRoom/common/message"
)

type UserProcess struct {
	//字段

}

//关联一个用户登录的方法
//写一个函数，完成登录校验
func (this *UserProcess) Login(userId int, userPwd string) (err error) {
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
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg() //mes 就是
	if err != nil {
		fmt.Println("ReadPkg er=", err)
		return
	}
	//将mes.Data 反序列化成LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)

	if loginResMes.Code == 200 {
		//初始化CurUser
		curUser.Conn = conn
		curUser.UserId = userId
		curUser.UserStatus = message.UserOnline

		//可以显示一下当前在线用户的列表，遍历loginResMes.UsersId
		fmt.Println("当前在线用户列表如下：")
		for _, v := range loginResMes.UsersId {
			//如果我们要求不显示自己在线，增加一个代码
			if v == userId {
				continue
			}
			fmt.Println("用户id\t", v)
			//完成客户端的onlineUsers 完成初始化
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Print("\n\n")
		//这里我们还需要在客户端启动一个协程
		//该协程保持和服务器端的通讯，如果服务器有数据推送给客户端
		//则接收并显示在客户端的终端
		go ServerProcessMes(conn)

		//显示登录成功后的菜单[循环显示]
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}
	return
}

//注册函数
func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
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
	mes.Type = message.RegisterMesType

	//3.创建一个RegisterMes 结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	//4.将registerMes序列化
	data, err := json.Marshal(registerMes)
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

	//创建一个transfer实例
	tf := &utils.Transfer{
		Conn: conn,
	}
	//发送data给服务器端
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("发送信息错误,err:", err)
	}

	mes, err = tf.ReadPkg() //mes 就是
	if err != nil {
		fmt.Println("ReadPkg er=", err)
		return
	}

	//将mes.Data 反序列化成RegisterResMes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)

	if registerResMes.Code == 200 {
		fmt.Println("注册成功，可以重新登录")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return
}
