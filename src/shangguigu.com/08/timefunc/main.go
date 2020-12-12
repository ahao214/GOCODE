package main

import(
	"fmt"
	"time"
)

//日期和时间的函数
func main(){
	//获取当前时间
	now:=time.Now()
	fmt.Println("当前时间：",now)

	//获取年月日时分秒
	fmt.Printf("年=%v\n",now.Year())
	fmt.Printf("月=%v\n",now.Month())
	fmt.Printf("月=%v\n",int(now.Month()))
	fmt.Printf("日=%v\n",now.Day())
	fmt.Printf("时=%v\n",now.Hour())
	fmt.Printf("分=%v\n",now.Minute())
	fmt.Printf("秒=%v\n",now.Second())


}
