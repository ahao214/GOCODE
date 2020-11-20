package main


import(
	"fmt"
	"strconv"

)

//strconv

func main(){
	// i:=int32(100)
	// ret:=int64(i)
	// fmt.Println(ret)

	// ret2:=fmt.Sprintf("%d",i)
	// fmt.Printf("%#v\n",ret2)

	str:="12345"
	//10表示十进制
	//64表示是64位
	ret1,err:=strconv.ParseInt(str,10,64)
	if err!=nil{
		fmt.Println("parseInt failed,err:%v\n",err)
		return
	}
	fmt.Printf("%#v %T\n",ret1,ret1)

}