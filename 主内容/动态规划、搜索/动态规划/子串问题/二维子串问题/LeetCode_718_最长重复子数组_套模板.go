package 二维子串问题

func findLength(A []int, B []int) int {
	/* 1. 明白dp[i][t]定义后，初始化dp数组 */
	dp := [1005][1005]int{} // 定义dp[i][t]为: 以A[i]、B[t]结尾的最长公共数组长度

	for i := 0; i < len(A); i++ {
		for t := 0; t < len(B); t++ {
			/* 2. dp[i][t]基础情况处理 (指: A[i] || B[t]前面没有元素时) */
			if i == 0 || t == 0{
				if A[i] == B[t]{
					dp[i][t] = 1
				}else{
					dp[i][t] = 0
				}
				continue
			}

			/* 3. 根据条件更新dp[i][t] */
			if A[i] == B[t] {
				dp[i][t] = dp[i-1][t-1] + 1
			} else {
				dp[i][t] = 0
			}
		}
	}

	/* 4. 获取最大/最小值 */
	ans := 0
	for i := 0; i < len(A); i++ {
		for t := 0; t < len(B); t++ {
			ans = max(ans, dp[i][t])
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

/*
	题目链接:
		https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray/		最长重复子数组
*/
/*
	总结
	1. 这题和最长相同子序列是不一样的，序列可以不连续，而这题是需要连续的。
	2. 官方还使用了二分法做...
*/
