package 一维子序列问题


// dp[i] 表示: 以nums[i]结尾的最长定差子序列长度 (注意，以下代码并不能AC题目链接的题，会超时, 时间复杂度O(n^2))
// 状态转移方程:
//		i == 0:							dp[i] = 1
// 		i >= 1:
// 			i >= 1 && nums[i] <= nums[t]:				dp[i] = max(dp[i], 1)		 		t ∈ [0, i)
//			i >= 1 && nums[i] - nums[t] == difference:	dp[i] = max(dp[i], dp[t]+1)		 	t ∈ [0, i)
func longestSubsequence(nums []int, difference int) int {
	if len(nums)==0{
		return 0
	}
	dp := make([]int, 100005)
	ans := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for t := 0; t < i; t++ {
			if nums[i] - nums[t] == difference {
				dp[i] = max(dp[i], dp[t]+1)
			}
		}
		ans = max(dp[i], ans)
	}
	return ans
}

// AC代码 (有dp数组) ( 时间复杂度O(n) )
func longestSubsequence(nums []int, difference int) int {
	if len(nums)==0{
		return 0
	}
	numToLength := make(map[int]int)	// 以num结尾的最长等差子序列长度
	dp := make([]int, 100005)
	ans := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = numToLength[nums[i]-difference] + 1
		numToLength[nums[i]] = dp[i]
		ans = max(dp[i], ans)
	}
	return ans
}

// 优化的AC代码 (消除dp数组) ( 时间复杂度O(n) )
func longestSubsequence(nums []int, difference int) int {
	if len(nums)==0{
		return 0
	}
	numToLength := make(map[int]int)
	ans := 0
	for i := 0; i < len(nums); i++ {
		numToLength[nums[i]] = numToLength[nums[i]-difference] + 1
		ans = max(numToLength[nums[i]], ans)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


/*
	题目链接:
		https://leetcode-cn.com/problems/longest-arithmetic-subsequence-of-given-difference/submissions/	最长定差子序列
*/

/*
	总结
	1. 感觉子序列的题dp都差不多..
	2. 为什么这题可以使用map优化呢？
			假如数组是 [2 4 9 6] difference == 2。
			对于2来说，它要形成的等差子序列的前一个数是确定的 0(2 - difference), 但是此时没有以0结尾的子序列，于是2自己成为1个子序列。
			之后走到了4，对于4来说，它要形成的等差子序列的前一个数也是确定的 2(4 - difference)，此时刚刚好有以2结尾的子序列，于是4的序列拓展了，变为了[2 4]
			....
		所以从上面可以看出，假如我们能确定子序列的前一个数是什么，那么我们就能使用map优化，而不用像最长上升子序列那样，
		必须扫描一次当前元素前面的所有元素。
*/
