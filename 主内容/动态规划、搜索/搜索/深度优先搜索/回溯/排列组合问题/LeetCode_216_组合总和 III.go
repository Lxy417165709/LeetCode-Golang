package 排列组合问题


// 回溯 + 外部变量 解决求数组组合问题 (数组中元素无重复且大于0, 不可重复选取，选定k个数)
var combinationSequence [][]int
func combinationSum3(k int, n int) [][]int {
	candidates := make([]int, 9)
	for i := 1; i <= 9; i++ {
		candidates[i-1] = i
	}
	combinationSequence = make([][]int, 0, 10)
	combinationSumExec(candidates, n, k, make([]int, 0, 10))
	return combinationSequence
}

func combinationSumExec(candidates []int, n int, k int, sequence []int) {

	if n == 0 && k == 0 {
		combinationSequence = append(combinationSequence, newSlice(sequence))
		return
	}
	if n == 0 || k == 0 {
		return
	}
	for i := 0; i < len(candidates); i++ {
		combinationSumExec(candidates[i+1:], n-candidates[i], k-1, append(sequence, candidates[i]))
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
		https://leetcode-cn.com/problems/combination-sum-iii/submissions/			组合求和3
*/

/*
	总结
	1. 该题题意是:
			在一个不可重复选取、元素大于0、元素不存在重复的数组中，选出k个数，使它们的值等于n，输出所有的组合。
*/
