package main

import (
	"encoding/json"
	"fmt"
)

//解析到interface
func main() {
	//假数据
	b := []byte(`{"age":"13","name":"枯藤","marry":false}`)
	//声明接口
	var i interface{}
	err := json.Unmarshal(b, &i)
	if err != nil {
		fmt.Println("json err:", err)
	}
	//自动转到map
	fmt.Println(i)
	//可以判断类型
	m := i.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case float64:
			fmt.Println(k, "是float64类型", vv)
		case string:
			fmt.Println(k, "是string类型", vv)
		default:
			fmt.Println("其他")
		}
	}
}
