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

//SQL注入

//SQL注入示例
func sqlInjectDemo(name string){
	strSql:=fmt.Sprintf("select id,name,age from user where name='%s'",name)

	fmt.Printf("SQL语句是:%s\n",strSql)
	var users []user
	err:=db.Select(&users,strSql)
	if err!=nil{
		fmt.Printf("exec failed,err:%v\n",err)
		return
	}
	
	for _,u:=range users{
		fmt.Printf("user:%#v\n",u)
	}
}


func main(){
	err:=initDB()
	if err!=nil{
		fmt.Printf("initDB failed,err:%#v\n",err)
		return
	}
	//几种SQL注入的例子
	sqlInjectDemo(" xxx' or 1=1 #")

}