package main

import(
	"fmt"
	"strings"
)


//14.最长公共前缀
func longestCommonPrefix(strs [3]string)string{
	if len(strs)<1{
		return ""
	}
	prefix:=strs[0]
	for _,k:=range strs{
		for strings.Index(k,prefix)!=0{
			if len(prefix)==0{
				return ""
			}
			prefix=prefix[:len(prefix)-1]
		}
	}
	return prefix
}


func main(){
	strs:=[3]string{"flo","flow","flower"}
	res:=longestCommonPrefix(strs)
	fmt.Println(res)
}