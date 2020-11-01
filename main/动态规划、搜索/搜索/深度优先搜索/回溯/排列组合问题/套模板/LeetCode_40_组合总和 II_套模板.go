package 套模板

import "sort"

var combinationSequence [][]int	// 结果集

// 返回结果集的函数
func combinationSum2(candidates []int, target int) [][]int {
	/* 1. 进行一些预处理 */
	combinationSequence = make([][]int, 0)
	sort.Ints(candidates)	// 有重复元素的组合问题就要排序。

	/* 2. 调用回溯函数 */
	combinationSumExec(candidates, target, make([]int, 0, 5))

	/* 5. 返回结果集 */
	return combinationSequence
}

// 回溯函数
func combinationSumExec(candidates []int, target int, sequence []int) {

	/* 3. 判断是否需要加入结果集以及进行剪枝 */
	if target < 0 {
		return
	}
	if target == 0 {
		combinationSequence = append(combinationSequence, newSlice(sequence))
		return
	}

	isVisited := make(map[int]bool)	// 题目有重复元素，所以需要这个结构
	for i := 0; i < len(candidates) && target >= candidates[i]; i++ {
		if isVisited[candidates[i]] == true {
			continue
		}
		isVisited[candidates[i]] = true

		/* 4. 继续调用回溯函数 */
		// 因为题目要求的是组合数且不能重复选取，所以下一层处理的是 candidates[i+1:]
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
