package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserId   int    `db:user_id`
	UserName string `db:username`
	Sex      string `db:sex`
	Email    string `db:email`
}

type Place struct {
	Country string `db:country`
	City    string `db:city`
	TelCode int    `db:telcode`
}

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed:", err)
		return
	}
	Db = database
	defer Db.Close()

}

func main() {
	res, err := Db.Exec("update person set username=? where user_id=?", "stu003", 1)
	if err != nil {
		fmt.Println("exec failed,", err)
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed,", err)

	}
	fmt.Println("update seccess:", row)

}
