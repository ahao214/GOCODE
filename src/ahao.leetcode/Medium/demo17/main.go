package main

import "fmt"

var (
	letterMap = []string{
		" ",    //0
		" ",    //1
		"abc",  //2
		"def",  //3
		"ghi",  //4
		"jkl",  //5
		"mno",  //6
		"pqrs", //7
		"tuv",  //8
		"wxyz", //9
	}
	res   = []string{}
	final = 0
)

//题目：17. 电话号码的字母组合
//说明：给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按任意顺序返回。
//输入：digits = "23"
//输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
//说明：数字 2-9 对应的电话的按键，数字1不对应任何字母
//解法一：DFS
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	res = []string{}
	findCombination(&digits, 0, "")
	return res
}

func findCombination(digits *string, index int, s string) {
	if index == len(*digits) {
		res = append(res, s)
		return
	}
	num := (*digits)[index]
	letter := letterMap[num-'0']
	for i := 0; i < len(letter); i++ {
		findCombination(digits, index+1, s+string(letter[i]))
	}
	return
}

//解法：回溯
var result []string
var dict = map[string][]string{
	"2": []string{"a", "b", "c"},
	"3": []string{"d", "e", "f"},
	"4": []string{"g", "h", "i"},
	"5": []string{"j", "k", "l"},
	"6": []string{"m", "n", "o"},
	"7": []string{"p", "q", "r", "s"},
	"8": []string{"t", "u", "v"},
	"9": []string{"w", "x", "y", "z"},
}

func letterCombinationsBT(digits string) []string {
	result = []string{}
	if digits == "" {
		return result
	}
	letterFunc("", digits)
	return result
}

func letterFunc(res string, digits string) {
	if digits == "" {
		result = append(result, res)
		return
	}

	k := digits[0:1]
	digits = digits[1:]
	for i := 0; i < len(dict[k]); i++ {
		res += dict[k][i]
		letterFunc(res, digits)
		res = res[0 : len(res)-1]
	}

}

func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))
	fmt.Println(letterCombinationsBT(digits))
}
