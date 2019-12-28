package main

import "fmt"

// dp[i][t] 表示 s[i:t+1]的最长回文子串长度
func LPD(s string) int {
	if len(s)==0{
		return 0
	}
	dp := [1050][1050]int{}
	for i := len(s) - 1; i >= 0; i-- {
		dp[i][i] = 1
		for t := i + 1; t < len(s); t++ {
			// 添加了这个条件: && dp[i+1][t-1] == t-i-1，因为要保证s[i-1:t]也是一个回文子串
			if s[i] == s[t] && dp[i+1][t-1] == t-i-1 {
				dp[i][t] = dp[i+1][t-1] + 2
			} else {
				dp[i][t] = max(dp[i+1][t],dp[i][t-1])
			}
		}
	}
	return dp[0][len(s)-1]
}

// dp[i][t] 表示 明确以s[i]、s[t]为左右边界的最长回文子串
func longestPalindrome(s string) string {
	dp := [1050][1050]int{}
	maxLength := 0
	ans :=""
	for i := len(s) - 1; i >= 0; i-- {
		dp[i][i] = 1
		if 1 > maxLength{
			maxLength = 1
			ans = s[i:i+1]
		}
		for t := i + 1; t < len(s); t++ {
			if s[i] == s[t] && dp[i+1][t-1] == t-i-1 {
				dp[i][t] = dp[i+1][t-1] + 2
			} else {
				dp[i][t] = 0
			}
			if dp[i][t] > maxLength{
				maxLength = dp[i][t]
				ans = s[i:t+1]
			}
		}
	}
	return ans
}

func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}
func main() {
	fmt.Println(LPD("1134567812344321214567897"))
}
