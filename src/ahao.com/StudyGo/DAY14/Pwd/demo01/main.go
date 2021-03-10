package main

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//加密解密-哈希算法
//使用场景
//对用户输入的密码进行加密
//用户登录时对用户的密码进行比对

func main() {
	userPwd := "123456"
	passwordbyte, err := GeneratePassword(userPwd)
	if err != nil {
		fmt.Println("密码加密出错")
	}
	fmt.Println(passwordbyte)
	passwordstring := string(passwordbyte)
	fmt.Println(passwordstring)
	//模拟这个字符串是从数据库读取出来的 值是12345678
	mysql_password := "$2a$10$I8WaWXgiBw8j/IBejb3t/.s5NoOYLvoQzL6mIM2g3TEu4z0HenzqK"
	isOk, _ := ValidatePassword(userPwd, mysql_password)
	if !isOk {
		fmt.Println("密码错误")
		return
	}
	fmt.Println(isOk)
}

//密码比对
func ValidatePassword(userPassword string, hashed string) (isOk bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, errors.New("密码对比错误")
	}
	return true, nil
}

//给密码就行加密操作
func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)

}
