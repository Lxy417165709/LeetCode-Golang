package 动态规划

import (
	"fmt"
	"sort"
)

// ------------------------ 字典树 + 记忆化搜索 ------------------------
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：4 MB, 在所有 Go 提交中击败了 100.00% 的用户
var canForm map[string]bool

func wordBreak(s string, wordDict []string) bool {
	trieRoot := NewTrie()
	canForm = make(map[string]bool)
	for i := 0; i < len(wordDict); i++ {
		trieRoot.InsertWord(wordDict[i])
	}
	return trieRoot.CanFormWord(s)
}

// ------------------ Trie ------------------
type Trie struct {
	nextNodes []*Trie
	isWord    bool
}

func NewTrie() *Trie {
	return &Trie{
		nextNodes: make([]*Trie, 26),
		isWord:    false,
	}
}

func (tr *Trie) InsertWord(word string) {
	for i := 0; i < len(word); i++ {
		if !tr.isCharExist(word[i]) {
			tr.createNextNode(word[i])
		}
		tr = tr.getNextNode(word[i])
	}
	tr.isWord = true
}

func (tr *Trie) CanFormWord(word string) bool {
	return tr.getCanFormWord(tr, word)
}

func (tr *Trie) getCanFormWord(root *Trie, word string) bool {
	if can, ok := canForm[word]; ok {
		return can
	}
	for i := 0; i < len(word); i++ {
		if !tr.isCharExist(word[i]) {
			canForm[word] = false
			return false
		}
		if tr.getNextNode(word[i]).isWord && root.getCanFormWord(root, word[i+1:]) {
			canForm[word] = true
			return true
		}
		tr = tr.getNextNode(word[i])
	}
	canForm[word] = tr.isWord
	return canForm[word]
}

func (tr *Trie) isCharExist(char byte) bool {
	return tr.nextNodes[char-'a'] != nil
}

func (tr *Trie) getNextNode(char byte) *Trie {
	return tr.nextNodes[char-'a']
}

func (tr *Trie) createNextNode(char byte) {
	tr.nextNodes[char-'a'] = NewTrie()
}

// ------------------------ 动态规划 ------------------------
// 执行用时：0 ms,   在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2.1 MB, 在所有 Go 提交中击败了 100.00% 的用户
//
// 时间复杂度：O(n^2 * m)
func wordBreak(s string, wordDict []string) bool {
	canForm := make([]bool, len(s)+1)
	canForm[0] = true
	for i := 1; i <= len(s); i++ {
		for t := 0; t < i; t++ {
			canForm[i] = canForm[i] || canForm[t] && isExistInDict(wordDict, s[t:i])
		}
	}
	return canForm[len(s)]
}
func isExistInDict(dict []string, ref string) bool {
	for i := 0; i < len(dict); i++ {
		if ref == dict[i] {
			return true
		}
	}
	return false
}

// ------------------------ 动态规划 + 二分查找 (优化 isExistInDict) ------------------------
// 执行用时：0 ms,   在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2.1 MB, 在所有 Go 提交中击败了 100.00% 的用户
//
// 时间复杂度：O(n^2 * log_m)
func wordBreak(s string, wordDict []string) bool {
	canForm := make([]bool, len(s)+1)
	canForm[0] = true
	sort.Strings(wordDict)
	for i := 1; i <= len(s); i++ {
		for t := 0; t < i; t++ {
			canForm[i] = canForm[i] || canForm[t] && isExistInDict(wordDict, s[t:i])
		}
	}
	return canForm[len(s)]
}
func isExistInDict(dict []string, ref string) bool {
	left, right := 0, len(dict)-1
	for left <= right {
		mid := (left + right) / 2
		if dict[mid] == ref {
			return true
		}
		if dict[mid] > ref {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

// ------------------------ 动态规划 + 哈希查找 ------------------------
// 执行用时：0 ms,   在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2.2 MB, 在所有 Go 提交中击败了 100.00% 的用户
//
// 时间复杂度：O(n^2)

func wordBreak(s string, wordDict []string) bool {
	isExistInDict := make(map[string]bool)
	for i := 0; i < len(wordDict); i++ {
		isExistInDict[wordDict[i]] = true
	}
	canForm := make([]bool, len(s)+1)
	canForm[0] = true
	for i := 1; i <= len(s); i++ {
		for t := 0; t < i; t++ {
			canForm[i] = canForm[i] || canForm[t] && isExistInDict[s[t:i]]
		}
	}
	return canForm[len(s)]
}

/*
	题目链接: https://leetcode-cn.com/problems/word-break/comments/
*/
