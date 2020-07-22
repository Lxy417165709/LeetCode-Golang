package 字符串

func maxPower(s string) int {
	if len(s) == 0 {
		return 0
	}
	curChar := s[0]
	curPower := 1
	maxPower := 1
	for i := 1; i < len(s); i++ {
		if s[i] != curChar {
			curPower = 1
			curChar = s[i]
			continue
		}
		curPower++
		maxPower = max(maxPower, curPower)
	}
	return maxPower
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
	题目链接: https://leetcode-cn.com/problems/consecutive-characters/submissions/
*/
