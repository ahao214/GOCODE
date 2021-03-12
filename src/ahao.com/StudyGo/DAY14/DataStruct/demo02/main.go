package common

import (
	"fmt"

	"github.com/student/1129/common"
)

//Product Product定义一个结构体
type Product struct {
	ID           int64  `json:"id" sql:"id" kuteng:"id"`
	ProductClass string `json:"ProductClass" sql:"ProductClass" kuteng:"ProductClass"`
	ProductName  string `json:"ProductName" sql:"productName" kuteng:"productName"`
	ProductNum   int64  `json:"ProductNum" sql:"productNum" kuteng:"productNum"`
	ProductImage string `json:"ProductImage" sql:"productImage" kuteng:"productImage"`
	ProductURL   string `json:"ProductUrl" sql:"productUrl"  kuteng:"productUrl"`
}

func main() {
	product := &Product{}
	//这块是表单提交的数据
	p.Ctx.Request().ParseForm()
	dec := common.NewDecoder(&common.DecoderOptions{TagName: "kuteng"})
	if err := dec.Decode(p.Ctx.Request().Form, product); err != nil {
		fmt.Println("绑定失败")
	}
}
