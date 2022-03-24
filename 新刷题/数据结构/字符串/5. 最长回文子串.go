package 字符串

import "fmt"

// longestPalindrome1 最长回文。 (朴素版)
func longestPalindrome2(s string) string {
	maxLength := 0
	result := ""
	dp := [1000][1000]int{}
	for i := len(s) - 1; i >= 0; i-- {
		for t := i; t <= len(s)-1; t++ {
			length := t - i + 1
			if length == 1 {
				dp[i][t] = 1
			} else if length == 2 {
				if s[i] == s[t] {
					dp[i][t] = 2
				} else {
					dp[i][t] = 0
				}
			} else {
				if s[i] == s[t] {
					if dp[i+1][t-1] != 0 {
						dp[i][t] = dp[i+1][t-1] + 2
					} else {
						dp[i][t] = 0
					}
				} else {
					dp[i][t] = 0
				}
			}
			if dp[i][t] > maxLength {
				result = s[i : t+1]
				maxLength = dp[i][t]
			}
		}
	}
	return result
}

// longestPalindrome2 最长回文。 (压缩空间)
func longestPalindrome1(s string) string {
	maxLength := 0
	result := ""
	dp := make([]int, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		for t := len(s) - 1; t >= i; t-- {
			length := t - i + 1
			if length == 1 {
				dp[t] = length
			} else if length == 2 {
				if s[i] == s[t] {
					dp[t] = length
				} else {
					dp[t] = 0
				}
			} else {
				if s[i] == s[t] {
					if dp[t-1] != 0 {
						dp[t] = length
					} else {
						dp[t] = 0
					}
				} else {
					dp[t] = 0
				}
			}
			if dp[t] > maxLength {
				result = s[i : t+1]
				maxLength = dp[t]
			}
		}
	}
	return result
}

// longestPalindrome3 最长回文。 (优美版+压缩空间)
func longestPalindrome3(s string) string {
	// 1. 初始化。
	result := ""
	maxLength := 0

	// 2. 动态规划。
	dp := [1000]bool{} // dp定义: dp[i][t] 表示 s[i:t+1] 是否为回文串。
	for i := len(s) - 1; i >= 0; i-- {
		for t := len(s) - 1; t >= 0; t-- {
			// 2.1 获取长度。
			length := t - i + 1

			// 2.2 动态规划。
			if length <= 2 {
				dp[t] = s[i] == s[t]
			} else {
				dp[t] = dp[t-1] && s[i] == s[t]
			}

			// 2.3 对比长度，更新最长回文子串。
			if dp[t] && length > maxLength {
				maxLength = length
				result = s[i : t+1]
			}
		}
	}

	// 3. 返回。
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
