package main

import (
	"fmt"
	"strconv"
)

//第299题：猜数字（Bulls and Cows）游戏
func main() {
	secret := "1807"
	guess := "7810"
	res := getHint(secret, guess)
	fmt.Println(res)
}

func getHint(secret string, guess string) string {
	a, b := 0, 0
	mapS, mapG := make([]int, 10), make([]int, 10)
	for i := range secret {
		//这里是获取对应数字的ASCII码
		tmp, charGuess := secret[i], guess[i]
		if tmp == guess[i] {
			a++
		} else {
			mapS[tmp-'0']++
			mapG[charGuess-'0']++
		}
	}
	for i := 0; i < 10; i++ {
		//找到重叠
		b += min(mapS[i], mapG[i])
	}
	//strconv.Itoa:整数转字符串
	return strconv.Itoa(a) + "A" + strconv.Itoa(b) + "B"

}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
