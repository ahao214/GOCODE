package mylogger


import(
	"os"
	"fmt"
	"time"
	"path"

)
//往文件里写日志相关代码

//文件日志结构体
type FileLogger struct{
	Level LogLevel
	filePath string 	//日志文件保存的路径
	fileName string 	//日志文件保存的文件名
	fileObj *os.File
	errFileObj *os.File
	maxFileSize int64
}


//NewFileLogger 构造函数
func NewFileLogger(levelStr,fp,fn string,maxSize int64) *FileLogger{
	logLevel,err:=parseLogLevel(levelStr)
	if err!=nil{
		panic(err)
	}
	fl:= &FileLogger{
		Level:logLevel,
		filePath:fp,
		fileName:fn,
		maxFileSize:maxSize,
	}
	err=fl.initFile()	//按照文件路径和文件名将文件打开
	if err!=nil{
		panic(err)
	}
	return fl
}

//根据指定的文件路径和文件名打开日志文件
func(f *FileLogger)initFile()(error){
	fullFileName:=path.Join(f.filePath,f.fileName)
	fileObj,err:=os.OpenFile(fullFileName,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err!=nil{
		fmt.Printf("open log file failed,err:%v\n",err)
		return err
	}

	errFileObj,err:=os.OpenFile(fullFileName+"err",os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err!=nil{
		fmt.Printf("open errLog file failed,err:%v\n",err)
		return err
	}

	//日志文件都已经打开了
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}

//切割日志文件
func(f *FileLogger) splitFile(file *os.File)(*os.File,error){	

			nowStr:=time.Now().Format("20060102150405000")
			fileInfo,err:=file.Stat()
			if err!=nil{
				fmt.Printf("open file info failed,err:%v\n",err)
				return nil,err
			}

			logName:=path.Join(f.filePath,fileInfo.Name())	//拿到当前日志文件的完整路径
			
			newLogName:=fmt.Sprintf("%s.bak%s",logName,nowStr)	//拼接一个日志文件备份的名字
			//1.关闭当前的日志文件
			file.Close()
			//2.备份一下 rename
			os.Rename(logName,newLogName)
			//3.打开一个新的日志文件
			fileObj,err:=os.OpenFile(logName,os.O_CREATE|os.O_APPEND|os.O_WRONLY,0644)
			if err!=nil{
				fmt.Printf("open new log file failed,err:%v\n",err)
				return nil,err
			}			
			//4.将打开的新的日志文件对象赋值给f.fileObj
			return fileObj,nil		
}

//记录日志的方法
func(f *FileLogger) log(lv LogLevel,format string,a...interface{}){
	if f.enable(lv){
		msg:=fmt.Sprintf(format,a...)
		now:=time.Now()
		funcName,fileName,lineNo:=getInfo(3)
		if f.checkSize(f.fileObj){
			newFile,err:=f.splitFile(f.fileObj)	//日志文件
			if err!=nil{
				return
			}
			f.fileObj=newFile

		}
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s\n",now.Format("2006-01-02 15:04:05"),getLogString(lv),fileName,funcName,lineNo,msg)
		if lv>=ERROR{
			if f.checkSize(f.errFileObj){
				newFile,err:=f.splitFile(f.errFileObj)	//日志文件
				if err!=nil{
					return
				}
				f.errFileObj=newFile
			}
			
			//如果要记录的日志级别大于等于error，则要在errlog里面再次记录一次
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d] %s\n",now.Format("2006-01-02 15:04:05"),getLogString(lv),fileName,funcName,lineNo,msg)
		}
	}
}

//判断是否需要记录该日志
func(f *FileLogger) enable(logLevel LogLevel)bool{
	return logLevel>=f.Level
}

//判断文件是否需要切割
func(f *FileLogger)checkSize(file *os.File)bool{
	fileInfo,err:=file.Stat()
	if err!=nil{
		fmt.Printf("get file info failed,err:%v\n",err)
		return false
	}
	//如果文件的大小大于等于文件最大值，开始切割，返回true
	return fileInfo.Size()>=f.maxFileSize
}

func(f *FileLogger) Debug(format string,a...interface{}){
	f.log(DEBUG,format,a...)
	
}

func(f *FileLogger) Info(format string,a...interface{}){
	f.log(INFO,format,a...)
}

func(f *FileLogger) Warning(format string,a...interface{}){
	f.log(WARNING,format,a...)
}

func(f *FileLogger) Error(format string,a...interface{}){
	f.log(ERROR,format,a...)
}


func(f *FileLogger) Fatal(format string,a...interface{}){
	f.log(FATAL,format,a...)
}

func(f *FileLogger) Trace(format string,a...interface{}){
	f.log(TRACE,format,a...)	
}


func(f *FileLogger)Close(){
	f.fileObj.Close()
	f.errFileObj.Close()
}