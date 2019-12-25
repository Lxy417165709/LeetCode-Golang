package 一维子序列问题

import "sort"

// 动态规划 解决最大子序和问题
// dp[i] 表示: 以nums[i]结尾的最大整除序列
// 状态转移方程:
//		i == 0: 								dp[i] = [ nums[i] ]
// 		i >= 1:
//			i >= 1 && nums[i]%nums[t] != 0:		dp[i] = max(dp[i], [ nums[i] ])						t ∈ [0, i)
// 			i >= 1 && nums[i]%nums[t] == 0:		dp[i] = max(dp[i], dp[t] append nums[i])			t ∈ [0, i)
//	tip:
//		[ nums[i] ] 表示只有一个元素nums[i]的数组(切片)
//		dp[t] append nums[i] 表示dp[t]数组后面追加上一个元素nums[i]
func largestDivisibleSubset(nums []int) []int {
	dp := make([][]int, len(nums))

	// 这里需要使数组递增，方便后续的处理，
	// 否则假如数组是 [1 8 4]，如果不进行排序，采用这个dp算法得出的结果是[1 4]
	sort.Ints(nums)
	ans := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		dp[i] = append(dp[i], nums[i])
		for t := 0; t < i; t++ {
			if nums[i]%nums[t] == 0 {
				// 注意这里要采用newSlice(dp[t])深拷贝，不然append函数会影响原来的dp[t]
				// 也因为使用了深拷贝，所以时空耗费多了
				dp[i] = max(dp[i], append(newSlice(dp[t]), nums[i]))
			}
		}
		ans = max(ans, dp[i])
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
