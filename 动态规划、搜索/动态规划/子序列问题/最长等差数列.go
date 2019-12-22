package main

// dp[hash(i,tolerance)]表示以nums[i]结尾且公差是tolerance的等差子序列长度 (方法1)
// 转移方程在下面方法2处
func longestArithSeqLength(nums []int) int {
	dp := make(map[int]int, 0)
	ans := 0
	for i := 0; i < len(nums); i++ {
		for t := 0; t < i; t++ {
			hash1, hash2 := hash(i, nums[i]-nums[t]), hash(t, nums[i]-nums[t])
			if dp[hash2] == 0 {
				dp[hash2] = 1
			}
			dp[hash1] = max(dp[hash1], dp[hash2]+1)
			ans = max(ans, dp[hash1])
		}
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func hash(a, b int) int {
	// 通过设置哈希，可以一定程度上降低时间，但是会增加错误率
	return ((b + 10005) << 20) | a // 哈希要设置正确.. 之前<<10太小了
}

// dp[i][tolerance] 表示以nums[i]结尾且公差是tolerance的等差子序列长度 (方法2)(这里还是使用了map)
// 状态转移方程:
//		dp[i][任意值] = 1
//		dp[i][tolerance] = max(dp[i][tolerance], dp[t][tolerance]+1)	t∈[0,i)
func longestArithSeqLength(nums []int) int {
	dp := make([]map[int]int, 0)
	ans := 0
	for i := 0; i < len(nums); i++ {
		dp = append(dp, make(map[int]int))
		for t := 0; t < i; t++ {
			tolerance := nums[i] - nums[t]
			// 这里赋值的原因是，dp[t][tolerance]的最小值就是1。 (如果可以把数据都初始化为1，那么这步就不需要)
			if dp[t][tolerance] == 0 {
				dp[t][tolerance] = 1
			}
			dp[i][tolerance] = max(dp[i][tolerance], dp[t][tolerance]+1)
			ans = max(ans, dp[i][tolerance])
		}
	}
	return ans
}

const (
	off = 10005 // 因为公差可能是负数，所以让公差加个偏移变为正数
)

// dp[i][tolerance + off] 表示以nums[i]结尾且公差是tolerance的等差子序列长度 (方法3)(这里使用数组)
// 提交的时候会超时，但是超时案例自己提交时可以运行，不会超时。
func longestArithSeqLength(nums []int) int {
	dp := [2005][20500]int{}
	ans := 0
	for i := 0; i < len(nums); i++ {
		for t := 0; t < i; t++ {
			tolerance := nums[i] - nums[t]
			if dp[t][tolerance+off] == 0 {
				dp[t][tolerance+off] = 1
			}
			dp[i][tolerance+off] = max(dp[i][tolerance+off], dp[t][tolerance+off]+1)
			ans = max(ans, dp[i][tolerance+off])
		}
	}
	return ans
}

/*
	题目链接:
		https://leetcode-cn.com/problems/longest-arithmetic-sequence/submissions/		最长等差数列
*/

/*
	总结
	1. 方法1之所以使用哈希，是为了把二维的信息转换为一维的数字。
	2. 方法2使用了map数组，所以不需要把二维信息转换为一维的数字。
	3. 方法3使用了二维数组，但是出现诡异的情况...
	4. 官方还有人使用二分法做出来的，但是我还不懂...
*/
