package main

import "sort"

// 回溯 + 外部变量 解决求数组组合问题 (数组中元素有重复且大于0, 不可重复选取)
var combinationSequence [][]int
func combinationSum2(candidates []int, target int) [][]int {
	combinationSequence = make([][]int, 0, 5)

	// 需要执行排序
	// 目的是 防止这种情况: [1 7 1], target == 8时, 如果没排序,那么结果会有 [1 7] 和 [7 1]。(原因是7后又出现了1)
	// 如果执行了排序,那么数组就变为了 1 1 7, 这样7的后面就不会出现先前的1了。
	sort.Ints(candidates) // ( 代码不同点① )
	combinationSumExec(candidates, target, make([]int, 0, 5))
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

	isVisited := make(map[int]bool)
	// 由于数组有序，那么可以添加 && target>=candidates[i] 条件进行优化    // ( 代码不同点② )
	for i := 0; i < len(candidates) && target >= candidates[i]; i++ {
		// 这里是防止该层选取重复元素，导致组合重复		 // ( 代码不同点③ )
		if isVisited[candidates[i]] == true {
			continue
		}
		isVisited[candidates[i]] = true

		// [i+1:] 表示一个元素只能选择1次	// ( 代码不同点④ )
		combinationSumExec(candidates[i+1:], target-candidates[i], append(sequence, candidates[i]))
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
		https://leetcode-cn.com/problems/combination-sum-ii/comments/			组合求和2
*/

/*
	总结
	1. 这题和 "求数组中等于目标值的组合.go" 很像，但代码需要修改一下。 (上面有4个代码修改点，已经标识出来了)
	2. 该题题意是:
			在一个不可重复选取、元素大于0、元素存在重复的数组中，选出一些数，使它们的值等于target，输出所有的组合。
			(可以和本文件夹下 "求数组中等于目标值的组合.go" 进行比较)
*/
