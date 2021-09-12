package main

import "fmt"

//242. 有效的字母异位词
func isAnagram(s string, t string) bool {
	hash := map[rune]int{}
	for _, value := range s {
		hash[value]++
	}
	for _, value := range t {
		hash[value]--
	}
	for _, value := range hash {
		if value != 0 {
			return false
		}
	}
	return true
}

func main() {
	isTrue := isAnagram("rat", "tar")
	fmt.Println(isTrue)
}
