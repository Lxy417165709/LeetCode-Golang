package 动态规划

const mod = 1000000007

func numTilings(N int) int {
	waysOfTiling := make([]int, N+3)
	waysOfTiling[0] = 1
	waysOfTiling[1] = 1
	waysOfTiling[2] = 2
	for i := 3; i <= N; i++ {
		waysOfTiling[i] = waysOfTiling[i-1] + waysOfTiling[i-2]
		for t := i - 3; t >= 0; t-- {
			waysOfTiling[i] += 2 * waysOfTiling[t]
		}
		waysOfTiling[i] %= mod
	}
	return waysOfTiling[N]
}

/*
	题目链接: https://leetcode-cn.com/problems/domino-and-tromino-tiling/submissions/
	总结
		1. 这题的思路和跳楼梯一样，不过要找出形成N阶骨牌的方式比找出形成N阶台阶的方式难一点。
		2. 上面的代码的时间复杂度是O(n^2)，可以把求和部分优化为O(1)，这样就可以使用O(n)时间复杂度AC这题了。
		3. 通过将DP表达式数学化简后，之后再进行遍历，也可以再O(n)时间内AC这题。
*/
