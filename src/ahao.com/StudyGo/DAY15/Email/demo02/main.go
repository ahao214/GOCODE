package main

import (
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
)

//实现抄送功能
func main() {
	e := email.NewEmail()
	//设置发送方的邮箱
	e.From = "dj<sss@qq.com>"
	//设置接收方的邮箱
	e.To = []string{"XXX@qq.com"}
	//设置抄送，如果抄送多人逗号隔开
	e.Cc = []string{"xx@qq.com", "ss@qq.com"}
	//设置密码抄送
	e.Bcc = []string{"xx@qq.com"}
	//设置主题
	e.Subject = "这是主题"
	//设置文件发送的内容
	e.Text = []byte("www.ahao.com")
	//设置服务器相关的配置
	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", "你的邮箱账号", "这块是你的授权码", "smtp.qq.com"))
	if err != nil {
		log.Fatal(err)
	}

}
