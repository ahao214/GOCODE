package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

)

//GO连接Mysql案例


var db *sql.DB	//是一个连接池对象

func initDB()(err error){
//数据库信息
	//第一个root表示数据库的用户名
	//第二个root表示数据库的密码
	//goday10表示数据库的名称
	dsn:="root:root@tcp(127.0.0.1:3306)/goday10"

	//连接数据库 mysql表示连接的是mysql数据库
	db,err=sql.Open("mysql",dsn)	//不会校验用户名和密码是否正确
	if err!=nil{	//dsn格式不正确的时候报错
		return
	}
	
	err=db.Ping()	//尝试连接数据库
	if err!=nil{
		return
	}
	return

}

type user struct{
	id int
	name string
	age int
}

func query(){


}


func insert(){


}

func main(){
	err:=initDB()

	if err!=nil{
		fmt.Printf("init DB failed,err:%v\n",err)		
	}
	fmt.Println("连接数据库成功")

	var u1 user
	//查询单条
	sqlStr:=`select id,name,age from user where id=1`
	//执行
	rowObj:=db.QueryRow(sqlStr)
	//获取结果
	rowObj.Scan(&u1.id,&u1.name,&u1.age)
	//打印结果
	fmt.Printf("u1:%#v\n",u1)
	
}