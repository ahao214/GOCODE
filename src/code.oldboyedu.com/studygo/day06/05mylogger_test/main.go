package main


import(
	"code.oldboyedu.com/studygo/day06/mylogger"
	"time"

)

var log mylogger.Logger		//声明一个全局的接口变量
//测试自己写的日志库

func main(){
	log=mylogger.NewConsoleLog("debug")
	log=mylogger.NewFileLogger("info","./","214.log",10*1024*1024)
	for{
		id:=10010
		name:="子韧"
		log.Debug("这是一条debug的日志")
		log.Info("这是一条info的日志")
		log.Warning("这是一条Warning的日志")
		log.Error("这是一条Error的日志,id:%d,name:%s",id,name)
		log.Fatal("这是一条Fatal的日志")
		log.Trace("这是一条Trace的日志")
		
		time.Sleep(time.Second * 2)

	}


}