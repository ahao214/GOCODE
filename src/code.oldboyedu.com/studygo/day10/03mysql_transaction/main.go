package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

)

//MySQL事务操作



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
	//设置数据库连接池最大连接数
	db.SetMaxOpenConns(10)
	return

}

type user struct{
	id int
	name string
	age int
}

func transaction(){
	//开启事务
	tx,err:=db.Begin()
	if err!=nil{
		fmt.Printf("begin failed,err:%v\n",err)
		return
	}

	//执行多条SQL语句
	strSql1:=`update user set age=age-2 where id=1`
	strSql2:=`update user set age=age+2 where id=1`

	_,err=tx.Exec(strSql1)
	if err!=nil{
		//回滚
		tx.Rollback()
		fmt.Println("执行SQL1出错了，要回滚")
		return
	}
	_,err=tx.Exec(strSql2)
	if err!=nil{
		//回滚
		tx.Rollback()
		fmt.Println("执行SQL2出错了，要回滚")
		return
	}

	//提交
	err=tx.Commit()
	if err!=nil{
		tx.Rollback()
		fmt.Println("提交出错了，要回滚")
		return 
	}

	fmt.Println("执行事务成功！")

}

func main(){
	err:=initDB()

	if err!=nil{
		fmt.Printf("init DB failed,err:%v\n",err)		
	}
	fmt.Println("连接数据库成功")

	
	
}