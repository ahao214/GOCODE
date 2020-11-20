package main

import (
	"fmt"
	"time"

)

func f1(){
	now:=time.Now()
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	//时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())




	//转换时间
	ret:=time.Unix(1564803667,0)
	fmt.Println(ret)

	//当前时间加24小时
	fmt.Println(now.Add(24*time.Hour))

	//定时器
	// timer:=time.Tick(time.Second)
	// for t:=range(timer){
	// 	fmt.Println(t)
	// }

	//时间格式化
	fmt.Println(now.Format("2006-01-02"))
	
	//Sub两个时间相减
	nextYear, err:=time.Parse("2006-01-02","2021-01-01")
	if err!=nil{
		fmt.Printf("parse time failed,err:%v\n",err)
		return
	}
	d:=now.Sub(nextYear)
	fmt.Println(d)
}

//时区
func f2(){

	now:=time.Now()
	fmt.Println(now)

	//明天的这个时间

	//根据字符串加载时区
	loc,err:=time.LoadLocation("Asia/Shanghai")
	if err!=nil{
		fmt.Printf("local loc failed,err:%v\n",err)
		return 
	}

}
func main(){
	f2()
	
}