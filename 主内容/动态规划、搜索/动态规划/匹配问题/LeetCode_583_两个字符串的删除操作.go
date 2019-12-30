package main


func minDistance(word1 string, word2 string) int {
	dp := [505][505]int{} // dp[i][t] 表示 word1[:i] 转换为 word2[:t] 所需操作数。

	// 这一步很重要
	dp[0][0] = 0
	for i := 1; i <= len(word1); i++ {
		dp[i][0] = dp[i-1][0] + 1
	}
	for i := 1; i <= len(word2); i++ {
		dp[0][i] = dp[0][i-1] + 1
	}
	for i := 1; i <= len(word1); i++ {
		for t := 1; t <= len(word2); t++ {
			if word1[i-1]==word2[t-1]{
				dp[i][t] = dp[i-1][t-1]
			}else{
				// 与编辑距离相比，就这里少了一个操作
				dp[i][t] = min(dp[i-1][t],dp[i][t-1])+1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}
func min(nums ...int) int{
	if len(nums)==1{
		return nums[0]
	}
	a,b := nums[0],min(nums[1:]...)
	if a>b{
		return b
	}
	return a
}

/*
	总结
	1. 这题和编辑距离是一样的，代码改一句话就可以了。
	2. 还可以用最长公共子序列解这个题。
*/
