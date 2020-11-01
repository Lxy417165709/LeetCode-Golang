package main

// 获取数组array中组成target的排列数 (数组元素>0，不存在重复元素，可重复选取)
// dp[i]: 表示数组组成i的排列数
func getCountOfArrangement(nums []int, target int) int {
	if target < 0 {
		return 0
	}
	dp := make([]int, target+1)
	dp[0] = 1

	// 总结点①
	for i := 1; i <= target; i++ {
		for t := 0; t < len(nums); t++ {
			if i >= nums[t] {
				dp[i] += dp[i-nums[t]]
			}
		}
	}
	return dp[target]
}

/*
	题目链接:
		https://leetcode-cn.com/problems/climbing-stairs/						爬楼梯
*/

/*
	总结
	1. 求组合数和求排列数的代码几乎一样，只是把 总结点① 处的内外层循环调换了。
*/
