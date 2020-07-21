package 前缀和

func minStartValue(nums []int) int {
	prefixSum := getPrefixSum(nums)
	minValue := min(prefixSum...)
	return max(-minValue, 0) + 1
}

func getPrefixSum(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	prefixSum := make([]int, len(nums))
	prefixSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
	}
	return prefixSum
}

func min(arr ...int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	a := arr[0]
	b := min(arr[1:]...)
	if a > b {
		return b
	}
	return a
}

func max(arr ...int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	a := arr[0]
	b := max(arr[1:]...)
	if a < b {
		return b
	}
	return a
}

/*
	题目链接: https://leetcode-cn.com/problems/minimum-value-to-get-positive-step-by-step-sum/
*/
