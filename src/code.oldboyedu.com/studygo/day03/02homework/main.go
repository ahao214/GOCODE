package main


import "fmt"

func main(){
	//判断字符串中汉字的数量
	// s1:="hello深圳"

	// var count  int
	// for _,c:=range s1{
	// 	if unicode.Is(unicode.Han,c){
	// 		count++
	// 	}
	// }
	// fmt.Println(count)

	// s2:="how do you do"
	// s3:=strings.Split(s2," ")
	// m1:=make(map[string]int,10)
	// for _,w:=range s3{
	// 	fmt.Println(w)
	// 	if _,ok:=m1[w]; !ok{
	// 		m1[w]=1
	// 	}else{
	// 		m1[w]++
	// 	}
	// }
	// for value,key:=range m1{
	// 	fmt.Println(value,key)
	// }


	//回文判断
	ss:="黄山叶落松落叶山黄"
	r:=make([]rune,0,len(ss))
	for _,c:=range ss{
		r=append(r,c)
	}
	for i:=0;i<len(r)/2;i++{
		if r[i]!=r[len(r)-1-i]{
			fmt.Println("不是回文")
			return 
		}
	}
	fmt.Println("是回文")
}
