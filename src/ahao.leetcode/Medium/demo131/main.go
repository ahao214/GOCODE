package main

//题目：131.分割回文串
//说明：给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串
//输入: "aab"
//输出:
//[
//  ["aa","b"],
//  ["a","a","b"]
//]

//方法一
func partition1(s string) [][]string {
	if s == "" {
		return [][]string{}
	}
	res, pal := [][]string{}, []string{}
	findPalindrome(s, 0, "", true, pal, &res)
	return res
}

func findPalindrome(str string, index int, s string, isPal bool, pal []string, res *[][]string) {
	if index == len(str) {
		if isPal {
			tmp := make([]string, len(pal))
			copy(tmp, pal)
			*res = append(*res, tmp)
		}
		return
	}
	if index == 0 {
		s = string(str[index])
		pal = append(pal, s)
		findPalindrome(str, index+1, s, isPal && isPalindrome(s), pal, res)
	} else {
		temp := pal[len(pal)-1]
		s = pal[len(pal)-1] + string(str[index])
		pal[len(pal)-1] = s
		findPalindrome(str, index+1, s, isPalindrome(s), pal, res)
		pal[len(pal)-1] = temp
		if isPalindrome(temp) {
			pal = append(apl, string(str[index]))
			findPalindrome(str, index+1, temp, isPal && isPalindrome(temp), pal, res)
			pal = pal[:len(pal)-1]
		}
	}
	return
}

func isPalindrome(s string) bool {
	slen := len(s)
	for i, j := 0, slen-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

//方法二
func partition2(s string) [][]string {
	result := [][]string{}
	size := len(s)
	if size == 0 {
		return result
	}
	current := make([]string, 0, size)
	dfs(s, 0, current, &result)
	return result
}

func dfs(s string, idx int, cur []string, result *[][]string) {
	start, end := idx, len(s)
	if start == end {
		temp := make([]string, len(cur))
		copy(temp, cur)
		*result = append(*result, temp)
		return
	}
	for i := start; i < end; i++ {
		if isPal(s, start, i) {
			dfs(s, i+1, append(cur, s[start:i+1]), result)
		}
	}

}
func isPal(str string, s, e int) bool {
	for s < e {
		if str[s] != str[e] {
			return false
		}
		s++
		e--
	}
	return true
}

func partition3(s string) [][]string {
	res := [][]string{}
	if len(s) == 0 {
		return res
	}
	cur := []string{}
	backtrack(s, 0, cur, &res)
	return res
}

func backtrack(s string, start int, cur []string, res *[][]string) {
	if start == len(s) {
		tmp := make([]string, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}
	for i := start; i < len(s); i++ {
		if isPalindrome3(s[start : i+1]) {
			cur = append(cur, string(s[start:i+1]))
			backtrack(s, i+1, cur, res)
			cur = cur[:len(cur)-1]
		}
	}
}

func isPalindrome3(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
