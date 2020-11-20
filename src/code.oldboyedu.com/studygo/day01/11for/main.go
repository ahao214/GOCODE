package main


import "fmt"

func main(){
	for i:=0;i<10;i++{
		fmt.Println(i)
	}

	s:="hello 詹姆斯"
	for i,v:=range s{
		fmt.Printf("%d-%c\n", i,v)
	}

	//哑元变量
	for _,v:=range s{
		fmt.Printf("%c\n",v)
	}


	for k:=1;k<10;k++{
		for m:=1;m<=k;m++{
			fmt.Printf("%d*%d=%d\t",m,k,m*k)
		}
		fmt.Println()
	}
}