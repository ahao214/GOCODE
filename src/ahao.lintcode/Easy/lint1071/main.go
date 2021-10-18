package main

import (
	"fmt"
	"sort"
)

/**
 * 1071 · 词典中最长的单词
 * @param words: a list of strings
 * @return: the longest word in words that can be built one character at a time by other words in words
 */
func longestWord(words []string) string {
	sort.Strings(words)
	mp := make(map[string]bool)
	var res string
	for _, word := range words {
		size := len(word)
		if size == 1 || mp[word[:size-1]] {
			if size > len(res) {
				res = word
			}
			mp[word] = true
		}
	}
	return res
}

func main() {
	word := []string{"a", "banana", "app", "appl", "ap", "apply", "apple"}
	fmt.Println(longestWord(word))
}
