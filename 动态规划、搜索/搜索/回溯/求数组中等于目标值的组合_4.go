package main

// 记忆化搜索 解决求数组排列问题  (数组中元素无重复且大于0, 可重复选取，总和为target)
var isHaving map[int]int
func combinationSum4(candidates []int, target int) int {
	isHaving = make(map[int]int)
	return combinationSumExec(candidates, target)
}

func combinationSumExec(candidates []int, target int) int {
	if target < 0 {
		return 0
	}
	if target == 0 {
		return 1
	}
	if x, ok := isHaving[target]; ok {
		return x
	}
	sum := 0
	for i := 0; i < len(candidates); i++ {
		// candidates[:] 之所以是这样，是因为此时我们还可以用之前用过的
		sum += combinationSumExec(candidates[:], target-candidates[i])
	}
	isHaving[target] = sum
	return isHaving[target]
}

/*
	题目链接:
		https://leetcode-cn.com/problems/combination-sum-iv/			组合求和4
		https://leetcode-cn.com/problems/coin-change-2/					零钱兑换2
*/

/*
	总结
	1. 这题并不是一个组合问题，而是一个排列问题，和题目零钱兑换2本质是一样的。
	2. 零钱兑换2部分我使用了动态规划，而这里使用了记忆化搜索，其实这只是同种思想(去除重叠子)的不同实现，
	   其中迭代实现的一般称为动态规划，而递归实现的一般叫记忆化搜索。
	3. 该题题意是:
			在一个可重复选取、元素大于0、元素不存在重复的数组中，选出一些数，使它们的值等于target，求所有排列数。
*/
