package 一维数组

// INF 无穷大。
const INF = 10000000000

// maxSubArray 获取最大连续子数组和。
func maxSubArray(nums []int) int {
	// 1. 异常情况。
	if len(nums) == 0 {
		panic("题目出错。")
	}

	// 2. 初始化。
	dp := -INF   // 以 nums[index] 结尾的最大子数组和。
	result := dp // 数组的最大子数组和。

	// 3. 动态规划。
	for index := range nums {
		dp = max(dp+nums[index], nums[index])
		result = max(result, dp)
	}

	// 4. 返回。
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
