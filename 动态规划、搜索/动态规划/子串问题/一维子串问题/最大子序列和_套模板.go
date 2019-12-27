package 一维子串问题

const inf = 100000000000 // 要定义的足够大，一般要大于int32

func maxSubArray(nums []int) int {
	/* 1. 搞清楚定义后，初始化dp数组 */
	dp := make([]int, len(nums)) // 定义dp[i]为: 以nums[i]为结尾的最大子串和

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			/* 2. dp[i]基础情况处理 (指: nums[i]前面没有元素时) */
			dp[i] = nums[i]
			continue
		}
		/* 3. 根据条件更新dp[i] */
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
	}
	ans := -inf
	for i := 0; i < len(dp); i++ {
		ans = max(ans, dp[i])
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
	题目链接: https://leetcode-cn.com/problems/maximum-subarray/submissions/
*/
