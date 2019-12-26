package 一维子串问题

/*
	给定一个未经排序的整数数组，找到最长且连续的的递增序列。
*/

// 原始dp (空间没有优化)
// dp[i] 表示: 以nums[i]结尾的最长连续递增序列长度
// 状态转移方程:
//		初始						: dp[i] = 1。
//		i > 0 && nums[i] > nums[i-1]: dp[i] = dp[i-1] + 1
func findLengthOfLCIS(nums []int) int {
	dp := make([]int, len(nums)+1)
	maxLength := 0
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = 1
		}
		maxLength = max(maxLength, dp[i])
	}
	return maxLength
}

// 滚动优化，将空间复杂度优化为O(1)，上面的空间复杂度是O(n)
func findLengthOfLCIS(nums []int) int {
	maxLength := 0
	dp := 0
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] > nums[i-1] {
			dp = dp + 1
		} else {
			dp = 1
		}
		maxLength = max(maxLength, dp)
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
	题目链接:
		https://leetcode-cn.com/problems/longest-continuous-increasing-subsequence/		最长连续递增序列
*/
/*
	总结
	1. 对于这个题，官方还有类似滑动窗口的解法。
*/