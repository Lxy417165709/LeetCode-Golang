package main

// dp[i][t] 表示从(0,0)走到(i,t)的不同路径数
// 状态转移方程:
//			i == 0 && t == 0: dp[i][t] = 1							(初始状态)
// 			i == 0 && t != 0: dp[i][t] = dp[0][t-1]					(表示只能从上边走到该点)
//			i != 0 && t == 0: dp[i][t] = dp[i-1][0]					(表示只能从左边走到该点)
//			i >= 1 && t >= 1: dp[i][t] = dp[i-1][t] + dp[i][t-1]	(表示只能从左边或上边走到该点)
func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	dp := [105][105]int{}
	dp[0][0] = 1
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0]
	}
	for t := 1; t < n; t++ {
		dp[0][t] = dp[0][t-1]
	}
	for i := 1; i < m; i++ {
		for t := 1; t < n; t++ {
			dp[i][t] = dp[i-1][t] + dp[i][t-1]
		}
	}
	return dp[m-1][n-1]
}

// 滚动数组优化空间复杂度，原来空间复杂度为O(mn)，优化后为O(n)
func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	dp := [105]int{}
	dp[0] = 1
	for t := 1; t < n; t++ {
		dp[t] = dp[t-1]
	}
	for i := 1; i < m; i++ {
		// dp[0] = dp[0]
		// 如果想知道为什么有这句废话，请看 "求左上角到右下角的不同路径_添加障碍物版.go" 的滚动数组优化部分
		for t := 1; t < n; t++ {
			dp[t] = dp[t] + dp[t-1]
		}
	}
	return dp[n-1]
}

/*
	题目链接:
		https://leetcode-cn.com/problems/unique-paths/		不同路径

*/

/*
	总结
	1. 注意题目的限制: 每次只能向下或者向右移动一步
	2. 其实这题还可以用排列组合AC。
*/
