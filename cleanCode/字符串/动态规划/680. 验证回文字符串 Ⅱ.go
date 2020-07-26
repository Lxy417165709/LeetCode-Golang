package 动态规划

func validPalindrome(s string) bool {
	return isValidPalindrome(s, 0, 1)
}

func isValidPalindrome(s string, nowTimesOfMismatching int, maxAllowingTimesOfMismatching int) bool {
	if nowTimesOfMismatching > maxAllowingTimesOfMismatching {
		return false
	}
	if len(s) <= 1 {
		return true
	}
	firstChar, lastChar := s[0], s[len(s)-1]
	if firstChar == lastChar {
		return isValidPalindrome(s[1:len(s)-1], nowTimesOfMismatching, maxAllowingTimesOfMismatching)
	} else {
		return isValidPalindrome(s[:len(s)-1], nowTimesOfMismatching+1, maxAllowingTimesOfMismatching) || isValidPalindrome(s[1:], nowTimesOfMismatching+1, maxAllowingTimesOfMismatching)
	}
}

/*
	题目链接: https://leetcode-cn.com/problems/valid-palindrome-ii/
	总结:
		1. 这道题我是用递归AC的，其实可以把它转换为DP。
*/
