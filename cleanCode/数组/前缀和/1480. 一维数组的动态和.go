package 前缀和

// ----------------------- 前缀和 -----------------------
func runningSum(nums []int) []int {
	return getPrefixSum(nums)
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

// ----------------------- 前缀和原地存储 -----------------------
// 执行用时：0 ms,   在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2.6 MB, 在所有 Go 提交中击败了 100.00% 的用户
func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}

/*
	题目链接: https://leetcode-cn.com/problems/running-sum-of-1d-array/submissions/
*/
