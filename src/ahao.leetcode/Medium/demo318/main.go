package main

//318. 最大单词长度乘积
func maxProduct(words []string) int {
	if words == nil || len(words) == 0 {
		return 0
	}
	length, value, maxProduct := len(words), make([]int, len(words)), 0
	for i := 0; i < length; i++ {
		tmp := words[i]
		value[i] = 0
		for j := 0; j < len(tmp); j++ {
			value[i] |= 1 << (tmp[j] - 'a')
		}
	}
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if (value[i]&value[j]) == 0 && (len(words[i])*len(words[j]) > maxProduct) {
				maxProduct = len(words[i]) * len(words[j])
			}
		}
	}
	return maxProduct
}

func maxProduct2(words []string) (ans int) {
	masks := map[int]int{}
	for _, word := range words {
		mask := 0
		for _, ch := range word {
			mask |= 1 << (ch - 'a')
		}
		if len(word) > masks[mask] {
			masks[mask] = len(word)
		}
	}
	for x, lenX := range masks {
		for y, lenY := range masks {
			if x&y == 0 && lenX*lenY > ans {
				ans = lenX * lenY
			}
		}
	}
	return
}

func maxProduct3(words []string) (ans int) {
	masks := make([]int, len(words))
	for i, word := range words {
		for _, ch := range word {
			masks[i] |= 1 << (ch - 'a')
		}
	}
	for i, x := range masks {
		for j, y := range masks[:i] {
			if x&y == 0 && len(words[i])*len(words[j]) > ans {
				ans = len(words[i]) * len(words[j])
			}
		}
	}
	return
}
