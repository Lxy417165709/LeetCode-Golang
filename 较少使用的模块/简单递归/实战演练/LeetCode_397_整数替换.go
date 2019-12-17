package main

// 简单递归
func integerReplacement(n int) int {
	return integerReplacementExec(n)
}
func integerReplacementExec(n int) int {
	if n == 1 {
		return 0
	}
	ans := 0
	if n&1 == 1 {
		a, b := integerReplacementExec(n+1), integerReplacementExec(n-1)
		ans = min(a, b) + 1
	} else {
		ans = integerReplacementExec(n>>1) + 1
	}
	return ans
}

// 记忆化搜索
func integerReplacement(n int) int {
	times = make(map[int]int)
	return integerReplacementExec(n)
}

var times map[int]int
func integerReplacementExec(n int) int {
	if x, ok := times[n]; ok {
		return x
	}
	if n == 1 {
		return 0
	}
	ans := 0
	if n&1 == 1 {
		a, b := integerReplacementExec(n+1), integerReplacementExec(n-1)
		ans = min(a, b) + 1
	} else {
		ans = integerReplacementExec(n>>1) + 1
	}
	times[n] = ans
	return times[n]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*
	题目链接:
		https://leetcode-cn.com/problems/integer-replacement/submissions/	整数划分
*/

/*
	总结
	1. 这题直接用递归就可以AC了，而记忆化搜索是为了防止重复求解。
*/
