package process

import (
	"fmt"
	"shangguigu.com/ChatRoom/client/model"
	"shangguigu.com/ChatRoom/common/message"
)

//客户端要维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var curUser model.CurUser //在用户登录成功后，完成对CurUser的初始化

//在客户端显示当前在线的用户
func outputOnlineUser() {
	//遍历onlineUsers
	fmt.Println("当前在线用户列表：")
	for id, _ := range onlineUsers {
		fmt.Println("用户id：\t", id)
	}
}

//处理返回的NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok { //原来没有
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user
	outputOnlineUser()
}
