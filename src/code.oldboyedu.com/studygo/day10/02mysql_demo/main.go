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
	//设置数据库连接池最大连接数
	db.SetMaxOpenConns(10)
	return

}

type user struct{
	id int
	name string
	age int
}

//查询单条数据
func queryOne(id int){
	var u1 user
	//查询单条
	sqlStr:=`select id,name,age from user where id=?`
	//执行
	rowObj:=db.QueryRow(sqlStr,id)
	//获取结果
	rowObj.Scan(&u1.id,&u1.name,&u1.age)
	//打印结果
	fmt.Printf("u1:%#v\n",u1)

}

//查询多条数据
func queryMore(n int){
	sqlStr:=`select id,name,age from user where id>?`
	rows,err:=db.Query(sqlStr,n)
	if err!=nil{
		fmt.Printf("exec %s query failed,err:%v\n",sqlStr,err)
		return
	}
	//一定要关闭
	defer rows.Close()
	//循环数据
	for rows.Next(){
		var u user
		err:=rows.Scan(&u.id,&u.name,&u.age)
		if err!=nil{
			fmt.Printf("scan failed,err:%v\n",err)
		}
		fmt.Printf("u:%#v\n",u)
	}
}

//插入
func insert(){
	sqlStr:=`insert into user(name,age) values("李二",25)`
	ret,err:=db.Exec(sqlStr)
	if err!=nil{
		fmt.Printf("insert failed,err:%v\n",err)
		return
	}
	//如果是插入数据的操作，能够拿到插入数据的的ID
	id,err:=ret.LastInsertId()
	if err!=nil{
		fmt.Printf("get id failed,err:%v\n",err)
		return
	}
	fmt.Println("id:",id)


}


//更新
func updateRow(newAge int,id int){
	sqlStr:=`update user set age=? where id=?`
	ret,err:=db.Exec(sqlStr,newAge,id)
	if err!=nil{
		fmt.Printf("update failed,err:%v\n",err)
		return
	}
	n,err:=ret.RowsAffected()
	if err!=nil{
		fmt.Printf("get id failed,err:%v\n",err)
		return
	}
	fmt.Printf("更新了%d行数据\n",n)

}


//删除
func deleteRow(id int){
	sqlStr:=`delete from user where id=?`
	ret,err:=db.Exec(sqlStr,id)
	if err!=nil{
		fmt.Printf("delete failed,err:%v\n",err)
		return
	}
	n,err:=ret.RowsAffected()
	if err!=nil{
		fmt.Printf("get id failed,err:%v\n",err)
		return
	}
	fmt.Printf("更新了%d行数据\n",n)

}


//预处理方式插入多条数据
func prepareInsert(){
	sqlStr:=`insert into user(name,age) values(?,?)`
	stmt,err:=db.Prepare(sqlStr)
	if err!=nil{
		fmt.Printf("prepare failed,err:%v\n",err)
		return
	}
	defer stmt.Close()
	//后续只需要拿到stmt去执行一些操作
	var m= map[string]int{
		"小王子":20,
		"公主":19,
		"天机老人":120,
		"天宝老道":200,
	}
	for k,v:=range m{
		stmt.Exec(k,v)
	}
}


func main(){
	err:=initDB()

	if err!=nil{
		fmt.Printf("init DB failed,err:%v\n",err)		
	}
	fmt.Println("连接数据库成功")

	
	
}