package main

import (
	"bytes"
	"errors"
	"fmt"
)

func getRelativePath(path1, path2 string) string {
	if path1 == "" || path2 == "" {
		defer errors.New("参数不合适")
	}
	arr1, arr2 := []rune(path1), []rune(path2)
	//用来指向两个路径中不同目录的起始路径
	diff1, diff2 := 0, 0
	len1, len2 := len(arr1), len(arr2)
	var relativePath bytes.Buffer
	for i, j := 0, 0; i < len1 && j < len2; {
		//如果目录相同则往后遍历
		if arr1[i] == arr2[j] {
			if arr1[i] == '/' {
				diff1, diff2 = i, j
			}
			i++
			j++
		} else {
			//不同的目录
			//把path1非公共部分的目录转换为../
			for diff1 < len1 {
				//碰到下一级目录
				if arr1[diff1] == '/' {
					relativePath.WriteString("../")
				}
				diff1++
			}
			//把path2的非公共部分的路径加到后面
			diff2++
			relativePath.WriteString(string(arr2[diff2:]))
			break
		}
	}
	return relativePath.String()
}

//求相对路径
func main() {
	path1 := "/qihoo/app/a/b/c/d/new.c"
	path2 := "qihoo/app/1/2/test.c"
	fmt.Println(getRelativePath(path1, path2))
}
