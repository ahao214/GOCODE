package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

)

//GO连接Mysql案例

func main(){
	//数据库信息
	//第一个root表示数据库的用户名
	//第二个root表示数据库的密码
	//goday10表示数据库的名称
	dsn:="root:root@tcp(127.0.0.1:3306)/goday10"

	//连接数据库 mysql表示连接的是mysql数据库
	db,err:=sql.Open("mysql",dsn)	//不会校验用户名和密码是否正确
	if err!=nil{	//dsn格式不正确的时候报错
		fmt.Printf("dsn:%s invalied,err:%v\n",dsn,err)
		return
	}
	
	err=db.Ping()	//尝试连接数据库
	if err!=nil{
		fmt.Printf("open %s failed,err:%v\n",dsn,err)
		return
	}

}