package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
)

func main(){
	resp,err:=http.Get("127.0.0.1:9090/xxx/")
	if err!=nil{
		fmt.Println("get url failed,err:%v\n",err)
		return
	}
	//从resp中把服务端返回的数据读出来
	// var data[]byte
	// resp.Body.Read()
	// resp.Body.Close()

	b,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		fmt.Printf("read resp.Body failed,err:%v\n",err)
		return
	}
	fmt.Println(string(b))
}