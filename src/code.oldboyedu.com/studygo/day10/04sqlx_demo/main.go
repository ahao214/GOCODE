package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"

)

//sqlx的操作



var db *sqlx.DB	//是一个连接池对象

func initDB()(err error){
//数据库信息
	//第一个root表示数据库的用户名
	//第二个root表示数据库的密码
	//goday10表示数据库的名称
	dsn:="root:root@tcp(127.0.0.1:3306)/goday10"

	//连接数据库 mysql表示连接的是mysql数据库
	db,err=sqlx.Connect("mysql",dsn)	
	if err!=nil{	
		return
	}
	

	//设置数据库连接池最大连接数
	db.SetMaxOpenConns(10)
	return

}

type user struct{
	ID int
	Name string
	Age int
}

func main(){
	err:=initDB()
	if err!=nil{
		fmt.Printf("initDB failed,err:%v\n",err)
		return
	}
	//查询一条数据
	strSql:=`select id,name,age from user where id=1`
	var u user
	db.Get(&u,strSql)
	fmt.Printf("u:%#v\n",u)


	//查询多条
	var userList=make([]user,0,10)
	strSql2:=`select id,name,age from user`
	err=db.Select(&userList,strSql2)
	if err!=nil{
		fmt.Printf("select failed,err:%v\n",err)
		return
	}
	fmt.Printf("userList:%#v\n",userList)
}