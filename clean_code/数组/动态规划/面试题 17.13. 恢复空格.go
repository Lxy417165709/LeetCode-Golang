package 动态规划

import "sort"


// 运行对比
//
// 		方法1: 632 ms	4.7 MB	Golang
// 		方法2: 236 ms	4.7 MB	Golang
// 		方法3: 20 ms	4.7 MB	Golang

// --------------------------- 方法1: 时间复杂度 O(n^2 * m) ---------------------------
func respace(dictionary []string, sentence string) int {
	minCountOfUnrecognizedChar := make([]int, len(sentence)+1)
	for end := 1; end <= len(sentence); end++ {
		minCountOfUnrecognizedChar[end] = minCountOfUnrecognizedChar[end-1] + 1
		for begin := 0; begin <= end; begin++ {
			wordInSentence := sentence[begin:end]
			if isInDictionary(dictionary, wordInSentence) {
				minCountOfUnrecognizedChar[end] = min(minCountOfUnrecognizedChar[end], minCountOfUnrecognizedChar[begin])
			} else {
				minCountOfUnrecognizedChar[end] = min(minCountOfUnrecognizedChar[end], minCountOfUnrecognizedChar[begin]+len(wordInSentence))
			}
		}
	}
	return minCountOfUnrecognizedChar[len(sentence)]
}

func isInDictionary(dictionary []string, word string) bool {
	for i := 0; i < len(dictionary); i++ {
		if dictionary[i] == word {
			return true
		}
	}
	return false
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// --------------------------- 方法2: 时间复杂度 O(n^2 * log_m) ---------------------------
func respace(dictionary []string, sentence string) int {
	minCountOfUnrecognizedChar := make([]int, len(sentence)+1)
	sort.Strings(dictionary)
	for end := 1; end <= len(sentence); end++ {
		minCountOfUnrecognizedChar[end] = minCountOfUnrecognizedChar[end-1] + 1
		for begin := 0; begin <= end; begin++ {
			wordInSentence := sentence[begin:end]
			if isInDictionary(dictionary, wordInSentence) {
				minCountOfUnrecognizedChar[end] = min(minCountOfUnrecognizedChar[end], minCountOfUnrecognizedChar[begin])
			} else {
				minCountOfUnrecognizedChar[end] = min(minCountOfUnrecognizedChar[end], minCountOfUnrecognizedChar[begin]+len(wordInSentence))
			}
		}
	}
	return minCountOfUnrecognizedChar[len(sentence)]
}

func isInDictionary(sortedDictionary []string, word string) bool {
	left, right := 0, len(sortedDictionary)-1
	for left <= right {
		mid := left + (right-left)/2
		if sortedDictionary[mid] == word {
			return true
		}
		if sortedDictionary[mid] > word {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// --------------------------- 方法3: 时间复杂度 O(nm) ---------------------------
func respace(dictionary []string, sentence string) int {
	minCountOfUnrecognizedChar := make([]int, len(sentence)+1)
	for end := 1; end <= len(sentence); end++ {
		minCountOfUnrecognizedChar[end] = minCountOfUnrecognizedChar[end-1] + 1
		for _, wordInDictionary := range dictionary {
			if end < len(wordInDictionary) {
				continue
			}
			wordInSentence := sentence[end-len(wordInDictionary) : end]
			if wordInDictionary == wordInSentence {
				minCountOfUnrecognizedChar[end] = min(minCountOfUnrecognizedChar[end], minCountOfUnrecognizedChar[end-len(wordInSentence)])
			}
		}
	}
	return minCountOfUnrecognizedChar[len(sentence)]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*
	题目链接: https://leetcode-cn.com/problems/re-space-lcci/submissions/
	总结
		1. 以上的时间复杂度是O(n^2 * m), m为字典的单词数
		2. 可以通过二分搜索，将时间复杂度优化为 O(n^2 log_m)
		3. 修改字符查找方式，将时间复杂度优化为 O(nm)
		4. 运行对比:
			方法1: 632 ms	4.7 MB	Golang
			方法2: 236 ms	4.7 MB	Golang
			方法3: 20 ms	4.7 MB	Golang

*/
