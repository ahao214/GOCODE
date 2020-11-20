package mylogger

import (
	"errors"
	"strings"
	"runtime"
	"path"
	"fmt"

)

type LogLevel uint64


//Logger 接口
type Logger interface{
	Debug(format string,a...interface{})
	Info(format string,a...interface{})
	Warning(format string,a...interface{})
	Error(format string,a...interface{})
	Fatal(format string,a...interface{})
	Trace(format string,a...interface{}
}

const (
	//定义日志级别 
	UNKNOW LogLevel = iota
	DEBUG 
	TRACE
	INFO
	WARNING
	ERROR
	FATAL

)


func parseLogLevel(s string) (LogLevel,error){
	s=strings.ToLower(s)
	switch s{
	case "debug":
		return DEBUG,nil
	case "trace":
		return TRACE,nil
	case "info":
		return INFO,nil
	case "warning":
		return WARNING,nil
	case "error":
		return ERROR,nil
	case "fatal":
		return FATAL,nil
	default:
		err:=errors.New("未知的错误级别")
		return UNKNOW,err
	}
}


func getLogString(lv LogLevel)string{
	switch lv{
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	}
	return "DEBUG"
}

func getInfo(skin int)(funcName,fileName string,lineNo int){
	pc,file,lineNo,ok:=runtime.Caller(skin)
	if !ok{		
		fmt.Println("runtime.Caller() failed\n")
		return
	}
	funcName=runtime.FuncForPC(pc).Name()
	fileName=path.Base(file)
	funcName=strings.Split(funcName,".")[1]

	return 
}


