package 字符串

// longestValidParentheses 贪心错误解法。
//  场景:
//       1. ()()(，此时会认为最长括号长度是4，实际上也是4。 (正确)
//       2. ()(()， 此时会认为最长括号长度是4，但实际上是2。 (错误)
func longestValidParentheses(s string) int {
	left := 0
	right := 0
	result := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			right++
			if right > left {
				right = 0
				left = 0
			} else {
				result = max(result, right*2)
			}
		} else {
			left++
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
