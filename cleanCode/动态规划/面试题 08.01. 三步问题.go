package 动态规划

func waysToStep(n int) int {
	const mod = 1000000007
	ways := make([]int, n+3)
	ways[0], ways[1], ways[2] = 1, 1, 2
	for i := 3; i <= n; i++ {
		ways[i] = ways[i-1] + ways[i-2] + ways[i-3]
		ways[i] %= mod
	}
	return ways[n]
}

/*
    题目链接: https://leetcode-cn.com/problems/three-steps-problem-lcci/submissions/
    总结
        1. 这题和青蛙跳楼梯的思路是一样的
		2. 可以通过状态压缩，将空间复杂度优化为O(1)
*/
