package main

//383. 赎金信
func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	//遍历杂志里面英文字母出现的次数
	m := make(map[byte]int)
	for i := 0; i < len(magazine); i++ {
		m[magazine[i]]++
	}
	//遍历信件里面的英文字母的个数
	for i := 0; i < len(ransomNote); i++ {
		if num, ok := m[ransomNote[i]]; ok {
			if num >= 1 {
				m[ransomNote[i]]--
			} else {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func canConstruct2(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	cnt := [26]int{}
	for _, ch := range magazine {
		cnt[ch-'a']++
	}
	for _, ch := range ransomNote {
		cnt[ch-'a']--
		if cnt[ch-'a'] < 0 {
			return false
		}
	}
	return true
}
