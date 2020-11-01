package 字符串

import "bytes"

func isPrefixOfWord(sentence string, searchWord string) int {
	sentence = " " + sentence
	searchWord = " " + searchWord
	index := bytes.Index(
		[]byte(sentence),
		[]byte(searchWord),
	)
	if index == -1 {
		return -1
	}
	return getCountOfBlank(sentence[:index+1])
}

func getCountOfBlank(str string) int {
	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			count++
		}
	}
	return count
}


/*
	题目链接: https://leetcode-cn.com/problems/check-if-a-word-occurs-as-a-prefix-of-any-word-in-a-sentence/
	总结:
		1. 这题感觉也可以用字典树。
*/
