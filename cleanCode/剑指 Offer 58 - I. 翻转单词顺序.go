package main

import (
	"bytes"
	"strings"
)

const space = ' '

func reverseWords(s string) string {
	stringHandler := NewStringHandler()
	return stringHandler.GetStringOfReverseWords(s)
}

type StringHandler struct {
	sample string
	words  []string
}

func NewStringHandler() *StringHandler {
	return &StringHandler{}
}

func (sh *StringHandler) GetStringOfReverseWords(s string) string {
	sh.sample = s
	sh.formWords()
	sh.reverseWords()
	return strings.Join(sh.words, string(space))
}

func (sh *StringHandler) formWords() {
	wordBuffer := bytes.Buffer{}
	str := sh.sample+string(space)
	for i := 0; i < len(str); i++ {
		if isCharSpace(str[i]) {
			if wordBuffer.Len() != 0 {
				sh.words = append(sh.words, wordBuffer.String())
				wordBuffer.Reset()
			}
			continue
		}
		wordBuffer.WriteByte(str[i])
	}
}

func (sh *StringHandler) reverseWords() {
	for i := 0; i < sh.getWordsSize()/2; i++ {
		sh.words[i], sh.words[sh.getWordsSize()-i-1] = sh.words[sh.getWordsSize()-i-1], sh.words[i]
	}
}

func (sh *StringHandler) getWordsSize() int {
	return len(sh.words)
}


func isCharSpace(char byte) bool {
	return char == space
}






/*
	题目链接: https://leetcode-cn.com/problems/fan-zhuan-dan-ci-shun-xu-lcof/
	总结：
		1. 其实可以通过面向对象的方式，将许多字符串的解题代码聚合到一个对象中！
		2. 一般来说，使用 buffer.Write 添加字符串比采用 '+' 快。
*/
