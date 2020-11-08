package 丑数

const inf = 1000000000000

func nthUglyNumber(n int) int {
	uglyFactors := []int{2, 3, 5}
	index := make(map[int]int)
	uglyNums := make([]int, n)
	uglyNums[0] = 1
	for _, factor := range uglyFactors {
		index[factor] = 0
	}
	for i := 1; i < n; i++ {
		// 1. 生成下一个丑数
		uglyNums[i] = inf
		for _, factor := range uglyFactors {
			uglyNums[i] = min(uglyNums[i], uglyNums[index[factor]]*factor)
		}

		// 2. 维护 index 含义
		for _, factor := range uglyFactors {
			if uglyNums[index[factor]]*factor == uglyNums[i] {
				index[factor]++
			}
		}
	}
	return uglyNums[n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

