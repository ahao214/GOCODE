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
		fmt.Println(user, "登录成功")
	}
	////如果用户的id=100 密码等于123456，则是正确，否则不合法
	//if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
	//	//合法
	//	loginResMes.Code = 200
	//
	//} else {
	//	//不合法
	//	loginResMes.Code = 500 //500 表示该用户不存在
	//	loginResMes.Error = "该用户不存在，请先注册，在使用"
	//}
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
