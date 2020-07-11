package 动态规划

func numWays(n int) int {
	const givenMaxStep = 100
	const givenMod = 1000000007
	jumpWays := getJumpWays(givenMod, givenMaxStep)
	return jumpWays[n]
}

func getJumpWays(mod int, maxStep int) []int {
	jumpWays := make([]int, maxStep+1)
	initialJumpWays(jumpWays)
	formJumpWays(jumpWays, mod)
	return jumpWays
}

func initialJumpWays(jumpWays []int) {
	jumpWays[0] = 1
	jumpWays[1] = 1
}

func formJumpWays(jumpWays []int, mod int) {
	for i := 2; i <= len(jumpWays)-1; i++ {
		jumpWays[i] = (jumpWays[i-1] + jumpWays[i-2]) % mod
	}
}

/*
	题目链接：https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/
*/
