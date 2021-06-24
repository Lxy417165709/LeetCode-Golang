package 数组动态规划

// 状态: 位置、手中有无数
// 操作: 不取、取
// 定义: 位于某状态的最长上升子序列长度
// dp[i][0] = max(dp[i-1][0],dp[i-1][1])
// dp[i][1] = max(1,dp[t][1]+1)
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	maxLength := make([][]int, n)
	for i := 0; i < n; i++ {
		maxLength[i] = make([]int, 2)
	}
	for i := 0; i < n; i++ {
		if i >= 1 {
			maxLength[i][0] = max(maxLength[i-1][0], maxLength[i-1][1])
		}
		maxLength[i][1] = 1
		for t := 0; t < i; t++ {
			if nums[i] > nums[t] {
				maxLength[i][1] = max(maxLength[i][1], maxLength[t][1]+1)
			}
		}
	}
	return max(maxLength[n-1][0], maxLength[n-1][1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


/*
	这个题还有O(nlogn)的解法，该解法采用二分查找。
*/