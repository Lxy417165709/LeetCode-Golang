package 双数组动态规划

// 状态: s[:i]、p[:t]
// 操作: 匹配
// 定义: 处于某状态下的匹配性
func isMatch(s string, p string) bool {
	n, m := len(s), len(p)
	dp := newMatrix(n+1, m+1)
	
	// 这题我们要考虑空串状态，这里我们用 string[:0] 表示
	dp[0][0] = true
	for t := 1; t < m; t += 2 {
		if p[t] == '*' {
			dp[0][t+1] = true
			continue
		}
		break
	}

	for i := 1; i <= n; i++ {
		for t := 1; t <= m; t++ {
			if s[i-1] == p[t-1] || p[t-1] == '.' {
				dp[i][t] = dp[i-1][t-1]
				continue
			}
			if p[t-1] == '*' {
				if s[i-1] == p[t-2] || p[t-2] == '.' {
					dp[i][t] = dp[i][t-1] || dp[i-1][t] || dp[i][t-2]
					continue
				}
				if t >= 2 {
					dp[i][t] = dp[i][t-2]
				} else {
					dp[i][t] = false
				}
				continue
			}
			dp[i][t] = false
		}
	}
	return dp[n][m]
}

func newMatrix(n, m int) [][]bool {
	matrix := make([][]bool, n)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]bool, m)
	}
	return matrix
}
