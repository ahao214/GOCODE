package Log

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

//实时读取文件内容
func main() {
	fileName := "./my.log"
	config := tail.Config{
		ReOpen:   true,                                 //重新打开
		Follow:   true,                                 //是否跟随
		Location: &tail.SeekInfo{Offset: 0, Whence: 2}, //从文件的哪个地方开始读
		Poll:     true,                                 //文件不存在不报错
	}
	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file failed,err:", err)
		return
	}
	var (
		line *tail.Line
		ok   bool
	)
	for {
		line, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("line:", line.Text)
	}

}