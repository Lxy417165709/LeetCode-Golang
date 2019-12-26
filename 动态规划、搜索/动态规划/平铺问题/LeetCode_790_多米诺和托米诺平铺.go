package 平铺问题

// 原始dp，没有化简的
// dp[i]表示: 2*i的面板有多少种方式可以平铺
// dp[i]  =  2*{dp[0]+...+dp[i-4] +dp[i-3]}+dp[i-2]+dp[i-1]
func numTilings(N int) int {
	mod := 1000000007
	dp := [1005]int{}
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= N; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % mod
		for t := 3; t <= i; t++ {
			dp[i] += 2 * dp[i-t]
			dp[i] %= mod
		}
	}
	return dp[N]
}

// 化简dp
// dp[i] = dp[i-1]*2 + dp[i-3]
func numTilings(N int) int {
	mod := 1000000007
	dp := [1005]int{}
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= N; i++ {
		dp[i] = dp[i-1]*2 + dp[i-3]
		dp[i] %= mod
	}
	return dp[N]
}

/*
	题目链接:
		https://leetcode-cn.com/problems/domino-and-tromino-tiling/		多米诺和托米诺骨牌
		https://www.51nod.com/Challenge/Problem.html#problemId=1031		骨牌覆盖
*/

/*
	总结
	1. 这题其实是骨牌覆盖的升级版
*/
