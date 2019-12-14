package main

const (
	size = 100
	inf  = 1000000000
)
// 动态规划 解决数组中能组成n的最少元素数
// dp[i]表示: 数组中能组成i的最少元素数	(dp数组初始化为inf)
// 状态转移方程:
//		i == 0:				dp[i] = 0
// 		i >= array[t]:		dp[i] = min(dp[i], dp[i - array[t]] + 1)		 t ∈ [0, len(array) - 1]
func numSquares(n int) int {
	// 平方数数组
	array := [size]int{}
	for i := 1; i <= size; i++ {
		array[i-1] = i * i
	}

	dp := [10500]int{}
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = inf
	}

	for t := 0; t < size; t++ {
		for i := array[t]; i <= n; i++ {
			dp[i] = min(dp[i], dp[i-array[t]]+1)
		}
	}
	/*
	上面是为了说明问题，所以开了个array数组，其实这个数组是可以省略的。
	参考下面的写法。
	for t := 0; t < size; t++ {
		number := t * t
		for i := number; i <= n; i++ {
			dp[i] = min(dp[i], dp[i - number]+1)
		}
	}
	*/
	return dp[n]
}

func min(a,b int) int{
	if a>b{
		return b
	}
	return a
}

/*
	题目链接:
		https://leetcode-cn.com/problems/perfect-squares/submissions/		完全平方数
		https://leetcode-cn.com/problems/coin-change/						零钱兑换
*/

/*
	总结
	1. 上面的代码解决的是LeetCode_279_完全平方数。
	2. 上面题目链接中的题目本质问题都是一样的: 求数组中能组成n的最少元素数(该数组中没有负数，且没有重复的值)

*/