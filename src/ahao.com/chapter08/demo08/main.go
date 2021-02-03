package main

import (
	"fmt"
	"reflect"
)

type TagType struct {
	flied1 bool   "是否有货"
	field2 string "商品名称"
	field3 int    "商品价格"
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}

func main() {
	tt := TagType{true, "Phone X", 5000}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}

}
