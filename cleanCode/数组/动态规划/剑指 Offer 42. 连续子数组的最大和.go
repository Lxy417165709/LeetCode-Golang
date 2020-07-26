package 动态规划

const INF = 1000000000000000

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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
	题目链接: https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/submissions/
	总结
		1. 可以采用类似原址哈希的思想，将dp的结果存放在原数组中。
*/
