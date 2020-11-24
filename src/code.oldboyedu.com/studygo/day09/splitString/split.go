package splitString

import(	
	"strings"
)

//切割字符串

func Split(str string,seq string)[]string{
	var ret=make([]string,0,strings.Count(str,seq)+1)
	index:=strings.Index(str,seq)
	for index>=0{
		ret=append(ret,str[:index])
		str=str[index+len(seq):]
		index=strings.Index(str,seq)
	}
	ret=append(ret,str)
	return ret

}



//斐波那契数列
func Fib(n int)int{
	if n<2{
		return n
	}
	return Fib(n-1)+Fib(n-2)
}