package 动态规划___哈希

func findSubstringInWraproundString(p string) int {
	if len(p) == 0 {
		return 0
	}

	// LCS: LongestContinuousString
	lengthOfLCSBackByLetter := make(map[byte]int)
	nowLengthOfLCS := 1
	lengthOfLCSBackByLetter[p[0]] = nowLengthOfLCS
	for i := 1; i < len(p); i++ {
		if isLetterBBackToLetterA(p[i-1], p[i]) {
			nowLengthOfLCS++
		} else {
			nowLengthOfLCS = 1
		}
		lengthOfLCSBackByLetter[p[i]] = max(lengthOfLCSBackByLetter[p[i]], nowLengthOfLCS)
	}
	return getValueSumOfMap(lengthOfLCSBackByLetter)
}

func isLetterBBackToLetterA(letterA, letterB byte) bool {
	if letterA == 'z' {
		return letterB == 'a'
	}
	return letterA+1 == letterB
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getValueSumOfMap(mp map[byte]int) int {
	sum := 0
	for _, value := range mp {
		sum += value
	}
	return sum
}


/*
	题目链接: https://leetcode-cn.com/problems/unique-substrings-in-wraparound-string/
	总结
		1. 这题看得一头雾水...之后看题解看了半天才理解了题意。
*/
