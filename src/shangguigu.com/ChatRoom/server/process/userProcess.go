package process

import (
	"encoding/json"
	"fmt"
	"net"
	"shangguigu.com/ChatRoom/common/message"
	"shangguigu.com/ChatRoom/server/model"
	"shangguigu.com/ChatRoom/server/utils"
)

type UserProcess struct {
	//字段
	Conn net.Conn
	//增加一个字段，表示该Conn是哪个用户
	UserId int
}

//编写通知所有在线的用户的方法
//userId 要通知其他在线用户，我上线了
func (this *UserProcess) NotifyOthersOnlineUser(userId int) {
	//遍历onlineUsers,然后一个一个发送NotifyUserStatusMes
	for id, up := range userMgr.onlineUsers {
		//过滤掉自己
		if id == userId {
			continue
		}
		//开始通知
		up.NotifyMeOnline(userId)
	}
}

func (this *UserProcess) NotifyMeOnline(userId int) {
	//组装我们的NotifyUserStatusMes
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType
	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline
	//将notifyUserStatusMes序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}
	//将序列化后的notifyUserStatusMes赋给mes.Data
	mes.Data = string(data)

	//对mes再次序列化,准备发送
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}
	//发送 创建Transfer实例
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline err:", err)
		return
	}
}

//编写一个ServerProcessLogin函数，专门处理登录
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//核心代码
	//1.先从mes中取出mes.data,并发序列号成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal failed,err=", err)
		return
	}

	//先声明一个 resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	//在声明一个loginResMes
	var loginResMes message.LoginResMes

	//我们需要到redis数据库完成验证
	//1.使用model.MyUserDao到redis验证
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)

	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 300
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误..."
		}
	} else {
		loginResMes.Code = 200
		//这里因为用户已经登录成功，我们就把该登录成功的用户放入到userMgr中
		//将登录成功的用户的userid 赋给this
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUsers(this)
		//通知其他的在线用户，我上线了
		this.NotifyOthersOnlineUser(loginMes.UserId)
		//将当前在线用户的id放入到loginResMes.UsersId
		//遍历userMgr.onlineUsers
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UsersId = append(loginResMes.UsersId, id)
		}
		fmt.Println(user, "登录成功")
	}

	//将loginResMes 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshall failed err=", err)
		return
	}

	//将data赋值给resMes
	resMes.Data = string(data)
	//对resMes 序列化,准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshall failed,err=", err)
		return
	}
	//发送data,我们将其
	//因为使用分层模式，我们先创建一个Transfer实例，然后读取
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return

}

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	//核心代码
	//1.先从mes中取出mes.data,
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal failed,err=", err)
		return
	}

	//先声明一个 resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	//在声明一个loginResMes
	var registerResMes message.RegisterResMes

	//我们需要到redis数据库完成注册
	//1.使用model.MyUserDao到redis验证
	err = model.MyUserDao.Register(registerMes.User)

	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误..."
		}
	} else {
		registerResMes.Code = 200
		fmt.Println("注册成功")
	}

	//将loginResMes 序列化
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshall failed err=", err)
		return
	}

	//将data赋值给resMes
	resMes.Data = string(data)
	//对resMes 序列化,准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshall failed,err=", err)
		return
	}
	//发送data,我们将其
	//因为使用分层模式，我们先创建一个Transfer实例，然后读取
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
