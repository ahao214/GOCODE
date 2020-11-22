package main

import(
	"net/http"
	"fmt"
	"io/ioutil"
)

func f1(w http.ResponseWriter, r *http.Request){
	//str:="hello lph!"
	//w.Write([]byte(str))
	b,err:=ioutil.ReadFile("./xx.txt")
	if err!=nil{
		w.Write([]byte(fmt.Sprintf("%v",err)))
	}
	w.Write(b)
}

func f2(w http.ResponseWriter,r *http.Request){
	// fmt.Println(r.URL)
	queryParam:=r.URL.Query()
	name:=queryParam.Get("name")
	age:=queryParam.Get("age")
	fmt.Println(name,age)
	
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))
	w.Write([]byte("ok"))
}


func main(){
	http.HandleFunc("/posts/",f1)
	http.HandleFunc("/xxx/",f2)
	http.ListenAndServe("127.0.0.1:9090",nil)
}
