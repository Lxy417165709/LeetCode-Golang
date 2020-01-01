package main

type Node struct {
	maxScore int
	ways     int
}

func pathsWithMaxScore(board []string) []int {
	dp := [105][105]Node{}
	m, n := len(board), len(board[0])
	mod := 1000000007
	dp[m-1][n-1] = Node{0, 1}

	for i := m - 1; i >= 0; i-- {
		for t := n - 1; t >= 0; t-- {
			ch := board[i][t]
			if ch == 'X' {
				dp[i][t] = Node{0, 0}
				continue
			}

			candidates := []Node{
				dp[i+1][t],
				dp[i][t+1],
				dp[i+1][t+1],
			}
			mx := 0
			// 获取得分最高的路径
			for j := 0; j <= 2; j++ {
				mx = max(mx, candidates[j].maxScore)
			}

			// 统计有多少条路径可以得分最高
			for j := 0; j <= 2; j++ {
				if candidates[j].maxScore == mx {
					dp[i][t].ways += candidates[j].ways
					dp[i][t].ways %= mod
				}
			}
			if ch != 'E' && ch != 'S' {
				dp[i][t].maxScore = mx + int(ch-'0')
			} else {
				dp[i][t].maxScore = mx
			}
		}
	}
	// 如果没有路径，那么依照题意，得分也要置0
	if dp[0][0].ways == 0 {
		dp[0][0].maxScore = 0
	}
	return []int{dp[0][0].maxScore, dp[0][0].ways}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
