package 未归类

func new21Game(N int, K int, W int) float64 {
	if K == 0 {
		return 1
	}
	p := make([]float64, K+W)
	preSum := make([]float64, K+W)
	p[0] = 1
	preSum[0] = 1
	for i := 1; i <= K+W-1; i++ {
		left := max(0, i-W)
		right := min(i-1, K-1)
		p[i] = (preSum[right] - preSum[left] + p[left]) / float64(W)
		preSum[i] = preSum[i-1] + p[i]
	}
	result := 0.0
	for i := K; i <= N && i < len(p); i++ {
		result += p[i]
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*
	总结
	1. 这题目用到了dp + 前缀和，dp想到了，但是没用前缀和导致提交超时。
*/
