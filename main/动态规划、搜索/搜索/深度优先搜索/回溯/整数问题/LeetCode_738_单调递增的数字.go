package 整数问题

var ans int

func monotoneIncreasingDigits(N int) int {
	ans = 0
	for i := 1; i <= 9; i++ {
		monotoneIncreasingDigitsExec(N, i)
	}
	return ans
}

func monotoneIncreasingDigitsExec(N int, nowNumber int) {
	if nowNumber > N {
		return
	}
	ans = max(ans, nowNumber)
	lastDigit := nowNumber % 10

	for i := lastDigit; i <= 9; i++ {
		monotoneIncreasingDigitsExec(N, nowNumber*10+i)
	}
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
/*
	总结
	1. 这题和顺次数很像，直接暴力回溯就OK了。
	2. 第一次没看清楚题意，以为是计算有多少个，于是稀里糊涂搞了大半天...
		看清楚后才发现是求最大值...
*/
