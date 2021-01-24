package process

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"shangguigu.com/ChatRoom/client/utils"
	"shangguigu.com/ChatRoom/common/message"
)

//显示登录成功后的界面...
func ShowMenu() {
	fmt.Println("---恭喜XXX登录成功---")
	fmt.Println("---1.显示在线用户列表---")
	fmt.Println("---2.发送消息---")
	fmt.Println("---3.信息列表---")
	fmt.Println("---4.退出系统---")
	fmt.Println("---请选择(1-4)---")
	var key int
	var content string
	smsProcess := &SmsProcess{}
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		//fmt.Println("显示用户列表")
		outputOnlineUser()
	case 2:
		fmt.Println("请输入你想对大家说的话")
		fmt.Scanf("%s\n", &content)
		smsProcess.SendGroupMes(content)
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("你选择了退出系统")
		os.Exit(0)
	default:
		fmt.Println("你输入的选项不正确，请重新输入!")

	}

}

//和服务器端保持通讯
func ServerProcessMes(conn net.Conn) {
	//创建一个transfer实例，不停的读取服务器发送的消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在读取服务器发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg() err:", err)
			return
		}
		//如果读取到了消息，又是下一步处理逻辑
		switch mes.Type {
		case message.NotifyUserStatusMesType: //通知有人上线了
			//1.取出NotifyUserStatusMes
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), notifyUserStatusMes)
			//2.把这个用户的信息，状态保存到客户map[int]user中
			updateUserStatus(&notifyUserStatusMes)
		//处理
		default:
			fmt.Println("服务器端返回了未知的消息类型")
		}
		//fmt.Printf("mes:%v\n", mes)
	}
}
