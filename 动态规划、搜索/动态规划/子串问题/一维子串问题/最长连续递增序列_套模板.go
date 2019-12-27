package 一维子串问题

/*
	给定一个未经排序的整数数组，找到最长且连续的的递增序列。
*/


func findLengthOfLCIS(nums []int) int {
	/* 1. 明白dp定义后，初始化dp数组 */
	dp := make([]int, len(nums)) // 定义dp[i]为: 以nums[i]为结尾的最长递增串长度
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			/* 2. dp[i]基础情况处理 (nums[i]前面没有元素时) */
			dp[i] = 1
			continue
		}
		/* 3. 根据条件更新dp[i] */
		if nums[i] > nums[i-1]{
			dp[i] = dp[i-1] + 1
		}else{
			dp[i] = 1
		}
	}
	maxLength := 0
	for i:=0;i<len(dp);i++{
		maxLength = max(maxLength,dp[i])
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