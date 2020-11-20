package main

import "fmt"

/*
分金币
名字里面包含e或者E，分一个金币
名字里面包含i或者I，分两个金币
名字里面包含o或者O，分三个金币
名字里面包含u或者U，分四个金币
*/
var (
	coins = 50
	users=[]string{
		"Matthew","Sarah","Augustus","Heidi","Emilie","Peter","Giana","Adriano","Aaron","Elizabeth",
	}
	
	distribution=make(map[string]int,len(users))
)

func dispatchCoin()(left int){
	for _,name:=range users{
		for _,c:=range name{
			switch c{
			case 'e','E':
				distribution[name]++
				coins--
			case 'i','I':
				distribution[name]+=2
				coins-=2
			case 'o','O':
				distribution[name]+=3
				coins-=3
			case 'u','U':
				distribution[name]+=4
				coins-=4
			}
		}
	}
	left = coins
	return
}

func main(){
	left:=dispatchCoin()
	fmt.Println("剩下金币个数：",left)

	for k,v:=range distribution{
		fmt.Printf("%s:%d\n",k,v)
	}

}

