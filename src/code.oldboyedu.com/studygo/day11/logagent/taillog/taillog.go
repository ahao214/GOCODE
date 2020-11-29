package taillog

import(
	"fmt"
	"time"
	_ "github.com/hpcloud/tail"
)

//专门从日志文件收集日志的模块

var(
	tailObj *tail.Tail
)

func Init(){
	config:=tail.Config{
		ReOpen:true,
		Follow:true,
		Location:&tail.SeekInfo{Offset:0,Whence:2},
		MustExist:false,
		Poll:true,
	}
	tailObj,err=tail.TailFile(fileName,config)
	if err!=nil{
		fmt.Println("tail file failed:err",err)
		return
	}
	return 


}


func ReadChan()<-chan *tail.Line{
	return tailObj.Lines
}

func ReadLog(){
	var(
		line *tail.Line
		ok bool
	)

	for{
		line,ok=<-tailObj.Lines
		if !ok{
			fmt.Printf("tail file close reopen,filename:%s\n",tailObj.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("line",line.Text)
	}
}


