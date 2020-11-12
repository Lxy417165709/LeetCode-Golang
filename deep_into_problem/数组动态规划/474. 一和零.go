package 数组动态规划

// 状态: 位置、0个数、1个数
// 操作: 不取当前的、取当前的
func findMaxForm(strs []string, m int, n int) int {
	length := len(strs)
	dp := newCube(length+1, m+1, n+1)
	for i := 1; i <= length; i++ {
		zero, one := getZeroOne(strs[i-1])
		for pm := 0; pm <= m; pm++ {
			for pn := 0; pn <= n; pn++ {
				dp[i][pm][pn] = max(dp[i][pm][pn], dp[i-1][pm][pn])
				if pm-zero >= 0 && pn-one >= 0 {
					dp[i][pm][pn] = max(dp[i][pm][pn], dp[i-1][pm-zero][pn-one]+1)
				}
			}
		}
	}
	return dp[length][m][n]
}

func getZeroOne(str string) (int, int) {
	zero := 0
	for _, ch := range str {
		if ch == '0' {
			zero++
		}
	}
	return zero, len(str) - zero
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func newCube(a, b, c int) [][][]int {
	cube := make([][][]int, a)
	for i := 0; i < len(cube); i++ {
		cube[i] = newMatrix(b, c)
	}
	return cube
}

func newMatrix(n, m int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, m)
	}
	return matrix
}
