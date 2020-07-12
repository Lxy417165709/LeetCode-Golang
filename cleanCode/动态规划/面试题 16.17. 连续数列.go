package 动态规划

// -------------------------- 分治法 -----------------------
const INF = 1000000000

func maxSubArray(nums []int) int {
	return getMaxSumOfSubArray(nums)
}

func getMaxSumOfSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	mid := len(nums) / 2
	return max(getMaxSumWithMidNum(nums), getMaxSumOfSubArray(nums[:mid]), getMaxSumOfSubArray(nums[mid:]))
}

func getMaxSumWithMidNum(nums []int) int {
	mid := len(nums) / 2
	return getMaxSumWithLastNum(nums[:mid]) + getMaxSumWithFirstNum(nums[mid:])
}

func getMaxSumWithLastNum(nums []int) int {
	maxSum := -INF
	curSum := 0
	for i := len(nums) - 1; i >= 0; i-- {
		curSum += nums[i]
		maxSum = max(maxSum, curSum)
	}
	return maxSum
}

func getMaxSumWithFirstNum(nums []int) int {
	maxSum := -INF
	curSum := 0
	for i := 0; i < len(nums); i++ {
		curSum += nums[i]
		maxSum = max(maxSum, curSum)
	}
	return maxSum
}

func max(nums ...int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	a, b := nums[0], max(nums[1:]...)
	if a > b {
		return a
	}
	return b
}

// -------------------------- 动态规划法 -----------------------

func maxSubArray(nums []int) int {
	return getMaxSumOfSubArray(nums)
}

func getMaxSumOfSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxSumOfSubArrayEndPreNum := make([]int, len(nums)+1)
	maxSumOfSubArrayEndPreNum[0] = -INF
	for i := 1; i <= len(nums); i++ {
		maxSumOfSubArrayEndPreNum[i] = max(maxSumOfSubArrayEndPreNum[i-1]+nums[i-1], nums[i-1])
	}
	return getMax(maxSumOfSubArrayEndPreNum)
}

func getMax(nums []int) int {
	result := -INF
	for i := 0; i < len(nums); i++ {
		result = max(nums[i], result)
	}
	return result
}

func max(nums ...int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	a, b := nums[0], max(nums[1:]...)
	if a > b {
		return a
	}
	return b
}

/*
	题目链接: https://leetcode-cn.com/problems/contiguous-sequence-lcci/comments/
	总结:
		1. 这题和 _剑指 Offer 42. 连续子数组的最大和_ 一样。
*/
