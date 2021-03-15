package main

import (
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"
)

//实现邮件附件的发送
func main() {
	e := email.NewEmail()
	//设置发送方的邮箱
	e.From = "dj <XXX@qq.com>"
	// 设置接收方的邮箱
	e.To = []string{"XXX@qq.com"}
	//设置主题
	e.Subject = "这是主题"
	//设置文件发送的内容
	e.HTML = []byte(`
    <h1><a href="http://www.topgoer.com/">go语言中文网站</a></h1>    
    `)
	//这块是设置附件
	e.AttachFile("./test.txt")
	//设置服务器相关的配置
	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", "你的邮箱账号", "这块是你的授权码", "smtp.qq.com"))
	if err != nil {
		log.Fatal(err)
	}
}
