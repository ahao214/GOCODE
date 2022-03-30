package main

//345. 反转字符串中的元音字母
func reverseVowels(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; {
		if !isVowel(b[i]) {
			i++
			continue
		}
		if !isVowel(b[j]) {
			j--
			continue
		}
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
	return string(b)
}

func isVowel(s byte) bool {
	return s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u' || s == 'A' ||
		s == 'E' || s == 'I' || s == 'O' || s == 'U'
}

func reverseVowels2(s string) string {
	sByte := []byte(s)
	if sByte == nil || len(sByte) <= 1 {
		return s
	}
	dic := make(map[byte]bool)
	dic['a'] = true
	dic['e'] = true
	dic['i'] = true
	dic['o'] = true
	dic['u'] = true
	dic['A'] = true
	dic['E'] = true
	dic['I'] = true
	dic['O'] = true
	dic['U'] = true
	left, right := 0, len(sByte)-1
	for left < right {
		_, leftOK := dic[sByte[left]]
		if !leftOK {
			left++
		}
		_, rightOK := dic[sByte[right]]
		if !rightOK {
			right--
		}
		if leftOK && rightOK {
			tmp := sByte[left]
			sByte[left] = sByte[right]
			sByte[right] = tmp
			left++
			right--
		}
	}
	return string(sByte)
}
