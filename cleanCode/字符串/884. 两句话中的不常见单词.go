package 字符串

import "strings"

func uncommonFromSentences(A string, B string) []string {
	countOfWord := make(map[string]int)
	allWords := append(getWords(A), getWords(B)...)
	for i := 0; i < len(allWords); i++ {
		countOfWord[allWords[i]]++
	}
	uncommonWords := make([]string, 0)
	for word, count := range countOfWord {
		if count == 1 {
			uncommonWords = append(uncommonWords, word)
		}
	}
	return uncommonWords
}

func getWords(s string) []string {
	return strings.Split(s, " ")
}

/*
	题目链接: https://leetcode-cn.com/problems/uncommon-words-from-two-sentences/
*/
