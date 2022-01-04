package main

//22. 括号生成
func generateParenthesis(n int) []string {
	res := []string{}
	if n == 0 {
		return res
	}
	backtrack(n, n, "", &res)
	return res
}

func backtrack(left int, right int, cur string, res *[]string) {
	if left < 0 || right < 0 || right < left {
		return
	}

	if left == 0 && right == 0 {
		*res = append(*res, cur)
		return
	}
	//左边
	cur = cur + "("
	backtrack(left-1, right, cur, res)
	cur = cur[:len(cur)-1]

	//右边
	cur = cur + ")"
	backtrack(left, right-1, cur, res)
	cur = cur[:len(cur)-1]
}
