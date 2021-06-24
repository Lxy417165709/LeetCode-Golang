package 双数组动态规划

// 状态: nums1[:i],nums2[:t]
// 操作: 都不取、取A取B、取B不取A、取A不取B
// 定义: 某状态下的最大点积
func maxDotProduct(nums1 []int, nums2 []int) int {
	n, m := len(nums1), len(nums2)
	if n == 0 || m == 0 {
		return 0
	}
	dp := newMatrix(n+1, m+1)
	result := -100000000
	for i := 1; i <= n; i++ {
		for t := 1; t <= m; t++ {
			result = max(result, dp[i-1][t-1]+nums1[i-1]*nums2[t-1])     // 这里是为了防止一个都不选的情况，因为题目不允许一个都不选，而dp[n][m]包括了一个都不选的情况
			dp[i][t] = max(dp[i][t], dp[i][t-1])                         // 不取nums2[t]
			dp[i][t] = max(dp[i][t], dp[i-1][t])                         // 不取nums1[i]
			dp[i][t] = max(dp[i][t], dp[i-1][t-1]+nums1[i-1]*nums2[t-1]) // 二者都取
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

func newMatrix(n, m int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, m)
	}
	return matrix
}



