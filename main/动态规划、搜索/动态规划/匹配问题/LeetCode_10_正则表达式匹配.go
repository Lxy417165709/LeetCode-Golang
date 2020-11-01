package 匹配问题

func isMatch(s string, p string) bool {
	dp := [55][55]bool{} // dp[i][t] 表示 s[:i],p[:t]是否匹配

	// 这一步很重要
	dp[0][0] = true
	for i := 1; i <= len(p); i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-2]
		} else {
			dp[0][i] = false
		}
	}

	for i := 1; i <= len(s); i++ {
		for t := 1; t <= len(p); t++ {
			if s[i-1] == p[t-1] || p[t-1] == '.' {
				dp[i][t] = dp[i-1][t-1]
			} else {
				if p[t-1] == '*' && t >= 2 {
					if p[t-2] == s[i-1] || p[t-2] == '.' {
						dp[i][t] = dp[i-1][t] || dp[i][t-1] || dp[i][t-2]
					} else {
						dp[i][t] = dp[i][t-2]
					}
				}
			}
		}
	}
	return dp[len(s)][len(p)]
}
/*
	总结
	1. 我觉得这些题目有模板！
*/
