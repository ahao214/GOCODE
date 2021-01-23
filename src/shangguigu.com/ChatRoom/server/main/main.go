package main

import (
	"fmt"
	"net"
	"shangguigu.com/ChatRoom/server/model"
	"time"
)

//func readPkg(conn net.Conn)(mes message.Message,err error)  {
//	buf := make([]byte, 8094)
//	fmt.Println("等待客户端发送的数据...")
//	_,err = conn.Read(buf[:4])
//	if err != nil {
//		//err=errors.New("read pkg header error")
//		return
//	}
//
//	//根据我读到的buf[:4]长度，转成一个uint32
//	var pkgLen uint32
//	pkgLen=binary.BigEndian.Uint32(buf[0:4])
//
//	//根据pkgLen读取消息内容
//	n,err:=	conn.Read(buf[:pkgLen])
//	if n!=int(pkgLen) || err!=nil{
//		//err=errors.New("read pkg body error")
//		return
//	}
//
//	//把pkgLen反序列化成->message.Message
//	err=json.Unmarshal(buf[:pkgLen],&mes)
//	if err!=nil{
//		fmt.Println("反序列化失败,err=",err)
//		return
//	}
//
//	return
//}

//func writePkg(conn net.Conn,data []byte)(err error){
//	//先发送一个长度给对方
//	var pkgLen uint32
//	pkgLen = uint32(len(data))
//	var buf [4]byte
//	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
//	//发送长度
//	n, err := conn.Write(buf[:4])
//	if n != 4 || err != nil {
//		fmt.Println("conn.Write buf err=", err)
//		return
//	}
//
//	// 发送data本身
//	n, err = conn.Write(data)
//	if n !=int(pkgLen)  || err != nil {
//		fmt.Println("conn.Write data err=", err)
//		return
//	}
//	return
//}

//编写一个ServerProcessLogin函数，专门处理登录
//func serverProcessLogin(conn net.Conn, mes *message.Message)(err error)  {
//	//核心代码
//	//1.先从mes中取出mes.data,并发序列号成LoginMes
//	var loginMes message.LoginMes
//	err=json.Unmarshal([]byte(mes.Data),&loginMes)
//	if err!=nil{
//		fmt.Println("json.Unmarshal failed,err=",err)
//		return
//	}
//
//	//先声明一个 resMes
//	var resMes message.Message
//	resMes.Type=message.LoginResMesType
//	//在声明一个loginResMes
//	var loginResMes message.LoginResMes
//
//	//如果用户的id=100 密码等于123456，则是正确，否则不合法
//	if loginMes.UserId==100 && loginMes.UserPwd=="123456"{
//		//合法
//		loginResMes.Code=200
//
//	}else {
//		//不合法
//		loginResMes.Code=500	//500 表示该用户不存在
//		loginResMes.Error="该用户不存在，请先注册，在使用"
//	}
//	//将loginResMes 序列化
//	data,err:=json.Marshal(loginResMes)
//	if err!=nil{
//		fmt.Println("json.Marshall failed err=",err)
//		return
//	}
//
//	//将data赋值给resMes
//	resMes.Data=string(data)
//	//对resMes 序列化,准备发送
//	data,err=json.Marshal(resMes)
//	if err!=nil{
//		fmt.Println("json.Marshall failed,err=",err)
//		return
//	}
//	//发送data,我们将其
//	err=writePkg(conn,data)
//	return
//
//}

//编写一个ServerProcessMes 函数
//功能：根据客户端发送消息种类不同，决定调用哪个函数来处理
//func serverProcessMes(conn net.Conn, mes *message.Message)(err error){
//	switch mes.Type {
//	case message.LoginMesType:
//		//处理登录
//		err=serverProcessLogin(conn,mes)
//		default:
//		fmt.Println("消息类型不存在，无法处理!")
//	}
//	return
//}

//处理和客户端的通讯
func processTo(conn net.Conn) {
	//这里需要延时关闭
	defer conn.Close()

	//这里调用总控，创建一个
	processor := &Processor{
		Conn: conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("客户端和服务器通讯协程错误=err", err)
		return
	}
}

//编写一个函数，完成对userDao的初始化
func initUserDao() {
	//这里的pool本身就是一个全局变量
	//这里需要注意一个初始化顺序
	//先initPool,然后在initUserDao
	model.MyUserDao = model.NewUserDao(pool)

}

func main() {
	//当服务器启动时，我们就初始化redis链接池
	initPool("localhost:6379", 16, 0, time.Second)
	initUserDao()
	//提示信息
	fmt.Println("服务器在8889端口监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.listen err=", err)
		return
	}
	//一旦监听成功，就等待客户端来链接服务器
	for {
		fmt.Println("等待客户端来链接服务器...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}
		//一旦链接成功，则启动一个协程和客户端保持通讯
		go processTo(conn)
	}
}
