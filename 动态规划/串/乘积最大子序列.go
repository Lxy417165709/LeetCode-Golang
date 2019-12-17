package main

import "fmt"

// 动态规划 解决乘积最大子序列
// dp[i][t] 表示:
// 		dp[i][0] 表示以nums[i]结尾的连续乘积最小子序列
// 		dp[i][1] 表示以nums[i]结尾的连续乘积最大子序列
// 状态转移方程:
// 			i == 0: dp[0][0] = dp[0][1] = nums[0]
//			i >= 1:
//				dp[i][0] = min(nums[i], dp[i-1][0]*nums[i], dp[i-1][1]*nums[i])
//				dp[i][1] = max(nums[i], dp[i-1][0]*nums[i], dp[i-1][1]*nums[i])
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := [100005][2]int{}
	dp[0][0] = nums[0]
	dp[0][1] = nums[0]
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i][0] = min(nums[i], dp[i-1][0]*nums[i], dp[i-1][1]*nums[i])
		dp[i][1] = max(nums[i], dp[i-1][0]*nums[i], dp[i-1][1]*nums[i])
		ans = max(ans, dp[i][1])
	}
	return ans
}

func max(arr ...int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	a, b := arr[0], max(arr[1:]...)
	if a > b {
		return a
	}
	return b
}

func min(arr ...int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	a, b := arr[0], min(arr[1:]...)
	if a > b {
		return b
	}
	return a
}

// 空间压缩优化，将空间复杂度优化为O(1)  (上面的空间复杂度是O(n))
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxPro, minPro, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		lastMaxPro,lastMinPro := maxPro,minPro
		/*
			为什么不直接这样写呢？
			minPro = min(nums[i], minPro*nums[i], maxPro*nums[i])
			maxPro = max(nums[i], minPro*nums[i], maxPro*nums[i])

			因为:
				第一句执行完后，minPro就被更新了，就不是上一层的minPro了，
				此时maxPro需要的是上一层的minPro，所以会导致错误。
		*/
		minPro = min(nums[i], lastMinPro*nums[i], lastMaxPro*nums[i])
		maxPro = max(nums[i], lastMinPro*nums[i], lastMaxPro*nums[i])
		ans = max(ans, maxPro)
	}
	return ans
}

/*
	题目链接:
		https://leetcode-cn.com/problems/longest-palindromic-substring/submissions/		乘积最大子序列
*/

/*
	总结
	1. 这题挺有意思的，因为之前的子串题目，dp[i]之前只有2个状态(自成一串、和前面dp[i-1]的组成一串)，
	   而这题目有3个 (自成一串、和前面dp[i-1][0]的组成一串、和前面dp[i-1][1]的组成一串)。
*/