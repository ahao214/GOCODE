package main

//771. 宝石与石头
func numJewelsInStones(jewels string, stones string) int {
	if stones == "" {
		return 0
	}
	dic := make(map[string]int)
	for _, s := range stones {
		dic[string(s)]++
	}
	res := 0
	for _, j := range jewels {
		v, ok := dic[string(j)]
		if ok {
			res += v
		}
	}
	return res
}
