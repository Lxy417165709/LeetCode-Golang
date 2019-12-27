package 一维子序列问题

// 动态规划 解决最长上升子序列问题
func lengthOfLIS(nums []int) int {
	/* 1. 搞清楚定义后，初始化dp数组 */
	dp := make([]int, len(nums))    // 这里定义dp[i]为: 以nums[i]为结尾的最长上升子序列长度

	for i := 0; i < len(nums); i++ {
		/* 2. dp[i]基础情况处理 (指子序列只有一个元素时) */
		dp[i] = 1
		for t := 0; t < i; t++ {
			/* 3. 判断nums[i]与nums[t]的关系，决定是否更新dp[i] */
			if nums[i] > nums[t] {
				/* 4. 更新dp[i] */
				dp[i] = max(dp[i], dp[t]+1)
			}
		}
	}
	ans := 0
	for i := 0; i < len(dp); i++ {
		ans = max(dp[i], ans)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
	题目链接: https://leetcode-cn.com/problems/longest-increasing-subsequence/submissions/
*/
