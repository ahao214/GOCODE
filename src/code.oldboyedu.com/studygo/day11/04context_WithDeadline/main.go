package main

import(
	"fmt"
	"time"
	"context"
)

//context.WithDeadline

func main(){
	d:=time.Now().Add(time.Millisecond*50)
	ctx,cancel:=context.WithDeadline(context.Background(),d)
	defer cancel()
	select{
	case <- time.After(time.Second*1):
		fmt.Println("Tom")
	case <- ctx.Done():
		fmt.Println(ctx.Err())


	}


}