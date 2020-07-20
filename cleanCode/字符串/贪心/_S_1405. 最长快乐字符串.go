package 贪心

func longestDiverseString(a int, b int, c int) string {
	countOfChar := map[byte]int{'a': a, 'b': b, 'c': c}
	candidateChars := []byte{'a', 'b', 'c'}
	result := ""
	for {
		bestChar := getBestAppendableChar(result, countOfChar, candidateChars)
		if bestChar == 0 {
			break
		}
		result = result + string(bestChar)
		countOfChar[bestChar]--
	}
	return result
}

func getBestAppendableChar(str string, countOfChar map[byte]int, candidateChars []byte) byte {
	sort.Slice(candidateChars, func(i, t int) bool {
		return countOfChar[candidateChars[i]] > countOfChar[candidateChars[t]]
	})
	if canCharAppend(str, candidateChars[0]) && countOfChar[candidateChars[0]] > 0 {
		return candidateChars[0]
	}
	if canCharAppend(str, candidateChars[1]) && countOfChar[candidateChars[1]] > 0 {
		return candidateChars[1]
	}
	return 0
}

func canCharAppend(str string, candidateChar byte) bool {
	if len(str) < 2 {
		return true
	}
	if candidateChar == str[len(str)-1] && candidateChar == str[len(str)-2] {
		return false
	}
	return true
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*
	题目链接: https://leetcode-cn.com/problems/longest-happy-string/submissions/
	总结:
		1. 更好的方式是把 char和其个数封装为一个结构体，再对结构体进行排序。
		2. 可以用堆对这题进行优化。
*/
