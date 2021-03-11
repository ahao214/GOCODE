package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
)

//安全散列算法SHA（Secure Hash Algorithm）是美国国家安全局 （NSA） 设计，
// 美国国家标准与技术研究院（NIST） 发布的一系列密码散列函数，
// 包括 SHA-1、SHA-224、SHA-256、SHA-384 和 SHA-512 等变体。
// 主要适用于数字签名标准（DigitalSignature Standard DSS）里面定义的
// 数字签名算法（Digital Signature Algorithm DSA）。
// SHA-1已经不是那边安全了，google和微软都已经弃用这个加密算法。
// 为此，我们使用热门的比特币使用过的算法SHA-256作为实例。
// 其它SHA算法，也可以按照这种模式进行使用。

//sha1
func main() {
	str := "www.51mh.com"
	//方法一
	data := []byte(str)
	has := sha1.Sum(data)
	shastr1 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	fmt.Println(shastr1)

	//方法二
	w := sha1.New()
	io.WriteString(w, str) //将str写入到w中
	bw := w.Sum(nil)       //w.Sum(nil)将w的hash转成[]byte格式
	// shastr2 := fmt.Sprintf("%x", bw)    //将 bw 转成字符串
	shastr2 := hex.EncodeToString(bw) //将 bw 转成字符串
	fmt.Println(shastr2)
}

//sha256
func shaone() string {
	str := "www.5lmh.com"

	w := sha256.New()
	io.WriteString(w, str) //将str写入到w中
	bw := w.Sum(nil)       //w.Sum(nil)将w的hash转成[]byte格式

	// shastr2 := fmt.Sprintf("%x", bw)    //将 bw 转成字符串
	shastr2 := hex.EncodeToString(bw) //将 bw 转成字符串
	//fmt.Println(shastr2)
	return shastr2
}

func shatwo() string {
	str := "www.5lmh.com"

	w := sha512.New()
	io.WriteString(w, str) //将str写入到w中
	bw := w.Sum(nil)       //w.Sum(nil)将w的hash转成[]byte格式

	// shastr2 := fmt.Sprintf("%x", bw)    //将 bw 转成字符串
	shastr2 := hex.EncodeToString(bw) //将 bw 转成字符串
	//fmt.Println(shastr2)
	return shastr2
}
