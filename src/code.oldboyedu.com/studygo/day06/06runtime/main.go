package main



import(
	"fmt"
	"runtime"
	"path"
)

//runtime.Caller()

func main(){
	pc,file,line,ok:=runtime.Caller(0)
	if !ok{
		fmt.Println("runtime.Caller() failed\n")
		return
	}
	funcName:=runtime.FuncForPC(pc)
	fmt.Println(funcName)
	fmt.Println(file)	//06runtime/main.go
	fmt.Println(path.Base(file))
	fmt.Println(line)	//13
}