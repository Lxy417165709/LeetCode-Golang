package 匹配问题

func isMatch(s string, p string) bool {
	dp := [1050][1050]bool{} // dp[i][t] 表示 s[:i],p[:t]是否匹配

	// 这一步很重要
	dp[0][0] = true
	for i := 1; i <= len(p); i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-1]
		} else {
			dp[0][i] = false
		}
	}

	for i := 1; i <= len(s); i++ {
		for t := 1; t <= len(p); t++ {
			if s[i-1] == p[t-1] || p[t-1] == '?' {
				dp[i][t] = dp[i-1][t-1]
			} else {
				if p[t-1] == '*' {
					for j := 0; j <= i && dp[i][t] == false; j++ {
						dp[i][t] = dp[i][t] || dp[j][t-1]
					}
				}
			}
		}

	}
	return dp[len(s)][len(p)]
}
/*
	总结
	1. 这题目和正则表达式匹配的差不多 0.0
	2.  执行用时 :188 ms, 在所有 golang 提交中击败了6.88% 的用户
		内存消耗 :4.3 MB, 在所有 golang 提交中击败了100.00%的用户
	3. 虽然DP慢了，但DP更有普遍性。
	4. 官方有使用贪心做的，比DP快。
*/