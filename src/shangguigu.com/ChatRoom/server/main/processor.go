package main

import (
	"fmt"
	"io"
	"net"
	"shangguigu.com/ChatRoom/common/message"
	"shangguigu.com/ChatRoom/server/process"
	"shangguigu.com/ChatRoom/server/utils"
)

//先创建一个processor的结构体
type Processor struct {
	Conn net.Conn
}

//编写一个ServerProcessMes 函数
//功能：根据客户端发送消息种类不同，决定调用哪个函数来处理
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	//是否能接受到客户端发送的群发消息

	switch mes.Type {
	case message.LoginMesType:
		//处理登录
		//创建一个UserProcess实例
		up := &process.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		up := &process.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		//创建一个SmsMesType实例，完成转发群聊消息
		smsProcess := &process.SmsProcess{}
		smsProcess.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在，无法处理!")
	}
	return
}

func (this *Processor) process2() (err error) {
	//读客户端发送的信息
	for {
		//这里我们将读取数据包，直接封装成一个函数readPkg(),返回一个message，err

		//创建一个Transfer实例
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端也退出...")
				return err
			} else {
				fmt.Println("readPkg error=", err)
				return err
			}
		}
		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}
