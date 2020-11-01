package 一维子序列问题

import "sort"

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)

	/* 1. 搞清楚定义后，初始化dp数组 */
	dp := make([][]int, len(nums))   // 这里定义dp[i]为: 以nums[i]为结尾的最大整除子集

	for i := 0; i < len(nums); i++ {
		/* 2. dp[i]基础情况处理 (指子序列只有一个元素时) */
		dp[i] = append(dp[i], nums[i])
		for t := 0; t < i; t++ {
			/* 3. 判断nums[i]与nums[t]的关系，决定是否更新dp[i] */
			if nums[i]%nums[t] == 0 {
				/* 4. 更新dp[i] */
				dp[i] = max(dp[i], append(newSlice(dp[t]), nums[i]))
			}
		}

	}

	ans := make([]int, 0)
	for i := 0; i < len(dp); i++ {
		ans = max(dp[i], ans)
	}
	return ans
}

// 返回长度最长的切片
func max(a, b []int) []int {
	if len(a) > len(b) {
		return a
	}
	return b
}

// 切片深拷贝
func newSlice(slice []int) []int {
	s := make([]int, len(slice))
	copy(s, slice)
	return s
}

/*
	题目链接:
		https://leetcode-cn.com/problems/largest-divisible-subset/submissions/		最大整除子集
*/

/*
	总结
	1. 这题和 "最长上升子序列" 的思路是一样的，小伙伴可以对比一下，发现异同。
*/
