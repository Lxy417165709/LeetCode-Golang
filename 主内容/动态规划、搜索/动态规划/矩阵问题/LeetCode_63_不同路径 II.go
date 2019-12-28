package main

// dp[i][t] 表示从(0,0)走到(i,t)的不同路径数
// 状态转移方程:
//			i == 0 && t == 0: dp[i][t] = obstacleGrid[0][0] ^ 1									(初始状态)
// 			i == 0 && t != 0: dp[i][t] = dp[i-1][0] * (obstacleGrid[i][0] ^ 1)					(表示只能从上边走到该点)
//			i != 0 && t == 0: dp[i][t] = dp[i-1][0] * (obstacleGrid[0][t] ^ 1)					(表示只能从左边走到该点)
//			i >= 1 && t >= 1: dp[i][t] = (dp[i-1][t] + dp[i][t-1]) * (obstacleGrid[i][t] ^ 1)	(表示只能从左边或上边走到该点)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	}
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := [105][105]int{}

	/*
	    这里使用位运算简化了if判断，如果使用if判断的话，那么应该是
		if obstacleGrid[0][0] == 1{
			dp[0][0] = 0
		}else{
			dp[0][0] = 1
		}

		(obstacleGrid[i][t] ^ 1) 其实是位运算的一个小技巧，因为 0 ^ 1 == 1，而 1 ^ 1 == 0。
	*/
	dp[0][0] = obstacleGrid[0][0] ^ 1

	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] * (obstacleGrid[i][0] ^ 1)
	}
	for t := 1; t < n; t++ {
		dp[0][t] = dp[0][t-1] * (obstacleGrid[0][t] ^ 1)
	}
	for i := 1; i < m; i++ {
		for t := 1; t < n; t++ {
			dp[i][t] = (dp[i-1][t] + dp[i][t-1]) * (obstacleGrid[i][t] ^ 1)
		}
	}
	return dp[m-1][n-1]
}

// 滚动数组优化空间复杂度，原来空间复杂度为O(mn)，优化后为O(n)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	}
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := [105]int{}
	dp[0] = obstacleGrid[0][0] ^ 1
	for t := 1; t < n; t++ {
		dp[t] = dp[t-1] * (obstacleGrid[0][t] ^ 1)
	}
	for i := 1; i < m; i++ {
		dp[0] = dp[0] * (obstacleGrid[i][0] ^ 1) // 这句很重要，因为我们在下面并没有处理dp[0],只是处理dp[1:]
		for t := 1; t < n; t++ {
			dp[t] = (dp[t] + dp[t-1]) * (obstacleGrid[i][t] ^ 1)
		}
	}
	return dp[n-1]
}

/*
	题目链接:
		https://leetcode-cn.com/problems/unique-paths-ii/		不同路径2

*/

/*
	总结
	1. 注意题目的限制: 每次只能向下或者向右移动一步且图中有障碍物，标1的是障碍物(不能走)，标0的是空地(可以走)。

*/
