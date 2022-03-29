package 二维数组

import "github.com/Lxy417165709/LeetCode-Golang/新刷题/util/math_util"

// --------------------------------------------- 1. 朴素动态规划(开始) ---------------------------------------------

// maximalSquare 获取矩阵中的正方形最大面积。
func maximalSquare(matrix [][]byte) int {
	// 1. 动态规划。
	var dp [301][301]int // dp[x][y] 为以 matrix[x-1][y-1] 为右下角的正方形的最大边长。
	for i := 0; i < len(matrix); i++ {
		for t := 0; t < len(matrix[i]); t++ {
			x := i + 1
			y := t + 1
			if matrix[i][t] == '0' {
				dp[x][y] = 0
			} else {
				length := math_util.Min(dp[x-1][y], dp[x][y-1])
				if matrix[i-length][t-length] == '1' {
					dp[x][y] = length + 1
				} else {
					dp[x][y] = length
				}
			}
		}
	}

	// 2. 获取最大面积。
	maxSquare := 0
	for i := 0; i < len(matrix); i++ {
		for t := 0; t < len(matrix[i]); t++ {
			x := i + 1
			y := t + 1
			maxSquare = math_util.Max(maxSquare, dp[x][y]*dp[x][y])
		}
	}

	// 3. 返回。
	return maxSquare
}

// --------------------------------------------- 1. 朴素动态规划(结束) ---------------------------------------------

// --------------------------------------------- 2. 优雅动态规划(开始) ---------------------------------------------

// maximalSquare 获取矩阵中的正方形最大面积。
func maximalSquare(matrix [][]byte) int {
	// 1. 动态规划。
	var dp [301][301]int // dp[x][y] 为以 matrix[x-1][y-1] 为右下角的正方形的最大边长。
	for i := 0; i < len(matrix); i++ {
		for t := 0; t < len(matrix[i]); t++ {
			x := i + 1
			y := t + 1
			if matrix[i][t] == '0' {
				dp[x][y] = 0
			} else {
				dp[x][y] = math_util.Min(dp[x-1][y], dp[x][y-1], dp[x-1][y-1]) + 1
			}
		}
	}

	// 2. 获取最大面积。
	maxSquare := 0
	for i := 0; i < len(matrix); i++ {
		for t := 0; t < len(matrix[i]); t++ {
			x := i + 1
			y := t + 1
			maxSquare = math_util.Max(maxSquare, dp[x][y]*dp[x][y])
		}
	}

	// 3. 返回。
	return maxSquare
}

// --------------------------------------------- 2. 优雅动态规划(结束) ---------------------------------------------
