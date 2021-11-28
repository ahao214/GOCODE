package main

//438. 找到字符串中所有字母异位词
func findAnagrams(s string, p string) []int {
	var freq [256]int
	result := []int{}
	if len(s) == 0 || len(s) < len(p) {
		return result
	}
	for i := 0; i < len(p); i++ {
		freq[p[i]-'a']++
	}
	left, right, count := 0, 0, len(p)

	for right < len(s) {
		if freq[s[right]-'a'] >= 1 {
			count--
		}
		freq[s[right]-'a']--
		right++
		if count == 0 {
			result = append(result, left)
		}
		if right-left == len(p) {
			if freq[s[left]-'a'] >= 0 {
				count++
			}
			freq[s[left]-'a']++
			left++
		}

	}
	return result
}

func findAnagrams2(s, p string) (ans []int) {
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return
	}

	var sCount, pCount [26]int
	for i, ch := range p {
		sCount[s[i]-'a']++
		pCount[ch-'a']++
	}
	if sCount == pCount {
		ans = append(ans, 0)
	}

	for i, ch := range s[:sLen-pLen] {
		sCount[ch-'a']--
		sCount[s[i+pLen]-'a']++
		if sCount == pCount {
			ans = append(ans, i+1)
		}
	}
	return
}
