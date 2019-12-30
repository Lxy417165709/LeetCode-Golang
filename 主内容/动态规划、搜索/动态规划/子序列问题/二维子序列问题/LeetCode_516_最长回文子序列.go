package 二维子序列问题

// 采用 dp 求解 最长回文子序列
// dp[i][t] 表示 s[i:t+1]的最长回文子序列
// 状态转移方程:
//		i >  t					: dp[i][t] = 0
//		i == t					: dp[i][t] = 1
//		s[i] == s[t]			: dp[i][t] = dp[i+1][t-1]+2
//      s[i] != s[t]			: dp[i][t] = max(dp[i+1][t], dp[i][t-1])
func longestPalindromeSubseq(s string) int {

	dp := [1050][1050]int{}
	for i := len(s) - 1; i >= 0; i-- {
		dp[i][i] = 1
		for t := i+1; t < len(s); t++ {
			if s[i] == s[t] {
				dp[i][t] = dp[i+1][t-1] + 2
			} else {
				dp[i][t] = max(dp[i+1][t], dp[i][t-1])
			}
		}
	}
	return dp[0][len(s)-1]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 采用 记忆化搜索 最长回文子序列
func longestPalindromeSubseq(s string) int {
	LPS = make(map[string]int)
	return longestPalindromeSubseqExec(s)
}

var LPS map[string]int

func longestPalindromeSubseqExec(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	if x, ok := LPS[s]; ok {
		return x
	}
	ans := 0
	if s[0] == s[len(s)-1] {
		ans = longestPalindromeSubseqExec(s[1:len(s)-1]) + 2
	} else {
		a := longestPalindromeSubseqExec(s[:len(s)-1])
		b := longestPalindromeSubseqExec(s[1:])
		ans = max(a, b)
	}
	LPS[s] = ans
	return ans

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
	题目链接:https://leetcode-cn.com/problems/longest-palindromic-subsequence/submissions/		最长回文子序列
*/

/*
	总结
	1. 由官方的结果来看，使用迭代的dp时空效率都高于使用递归的记忆化搜索。
			(dp: 44ms 10.8MB) (记忆化搜索: 376ms 59.7MB)
	2. 从代码的简洁度来看，dp的比记忆化搜索的简洁。
	3. 从设计难度上看(个人感觉)，记忆化搜索比dp容易设计、编码。
*/