package common

import (
	"fmt"

	"github.com/student/1129/common"
)

//Product Product定义一个结构体
type Product struct {
	ID           int64  `json:"id" sql:"id"`
	ProductClass string `json:"ProductClass" sql:"ProductClass"`
	ProductName  string `json:"ProductName" sql:"productName"`
	ProductNum   int64  `json:"ProductNum" sql:"productNum"`
	ProductImage string `json:"ProductImage" sql:"productImage"`
	ProductURL   string `json:"ProductUrl" sql:"productUrl" `
}

func main() {
	//这块是模拟mysql获取单条的数据反射到结构体
	data := map[string]string{"id": "1", "ProductClass": "blog", "productName": "5lmh.com", "productNum": "40", "productImage": "http://www.5lmh.com/", "productUrl": "http://www.5lmh.com/"}
	productResult := &Product{}
	common.DataToStructByTagSql(data, productResult)
	fmt.Println(*productResult)
	//这块是模拟mysql获取所有的数据反射到结构体
	Alldata := []map[string]string{
		{"id": "1", "ProductClass": "blog", "productName": "5lmh.com", "productNum": "40", "productImage": "http://www.5lmh.com/", "productUrl": "http://www.5lmh.com/"},
		{"id": "2", "ProductClass": "blog", "productName": "5lmh.com", "productNum": "40", "productImage": "http://www.5lmh.com/", "productUrl": "http://www.5lmh.com/"},
	}
	var productArray []*Product
	for _, v := range Alldata {
		Allproduct := &Product{}
		common.DataToStructByTagSql(v, Allproduct)
		productArray = append(productArray, Allproduct)
	}
	for _, vv := range productArray {
		fmt.Println(vv)
	}
}
