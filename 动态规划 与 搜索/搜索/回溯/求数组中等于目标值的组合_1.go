package main

// 回溯 + 外部变量 解决求数组组合问题 (数组中元素无重复且大于0,可重复选取)
var combinationSequence [][]int
func combinationSum(candidates []int, target int) [][]int {
	// 容量定义大点，这样可以避免扩容时产生额外时空花费
	combinationSequence = make([][]int, 0, 100)
	combinationSumExec(candidates, target, make([]int, 0, 100))
	return combinationSequence
}

func combinationSumExec(candidates []int, target int, sequence []int) {
	if target < 0 {
		return
	}
	if target == 0 {
		combinationSequence = append(combinationSequence, newSlice(sequence))
		return
	}

	for i := 0; i < len(candidates); i++ {
		// [i:] 是为了防止再选取之前的元素，导致出现重复组合。(元素可以重复选取)
		// [i+1:] 表示一个元素只能选择1次
		combinationSumExec(candidates[i:], target-candidates[i], append(sequence, candidates[i]))
	}
}

// 深拷贝
func newSlice(oldSlice []int) []int {
	slice := make([]int, len(oldSlice))
	copy(slice, oldSlice)
	return slice
}

/*
	题目链接:
		https://leetcode-cn.com/problems/combination-sum/			组合求和
		https://leetcode-cn.com/problems/coin-change-2/				零钱兑换2
*/

/*
	总结
	1. 切片可以初始化容量，这样可以避免当切片长度大于容量时，切片扩容导致的额外代价。
	2. 该题和零钱兑换2几乎一样，只是零钱兑换2求的是种数，而这里需要输出所有的组合。 (零钱兑换那使用了动态规划，这里使用了回溯)
	3. 该题题意是: 在一个可重复选取、元素大于0的数组中，选出一些数，使它们的值等于target，输出所有的组合。
*/
