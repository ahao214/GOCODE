package model

import (
	"net"
	"shangguigu.com/ChatRoom/common/message"
)

//因为在客户端，我们很多地方会使用到curUser
type CurUser struct {
	Conn net.Conn
	message.User
}
