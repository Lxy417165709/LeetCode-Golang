package 字符串

var parenthesesPairs []string

func generateParenthesis(n int) []string {
	parenthesesPairs = make([]string, 0)
	formParenthesesPairs(n, n, []byte{})
	return parenthesesPairs
}

func formParenthesesPairs(countOfLeftParentheses, countOfRightParentheses int, curParenthesesPair []byte) {
	if countOfLeftParentheses > countOfRightParentheses {
		return
	}
	if countOfLeftParentheses == 0 && countOfRightParentheses == 0 {
		parenthesesPairs = append(parenthesesPairs, string(curParenthesesPair))
		return
	}
	if countOfLeftParentheses < 0 || countOfRightParentheses < 0 {
		return
	}
	formParenthesesPairs(countOfLeftParentheses-1, countOfRightParentheses, append(curParenthesesPair, '('))
	formParenthesesPairs(countOfLeftParentheses, countOfRightParentheses-1, append(curParenthesesPair, ')'))
}

/*
	题目链接: https://leetcode-cn.com/problems/bracket-lcci/
*/
