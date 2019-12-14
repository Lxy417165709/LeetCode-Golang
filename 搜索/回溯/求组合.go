package main

// 回溯 + 外部变量 解决求数组组合问题 (该数组没有重复元素时)
var combineSequence [][]int
func combine(n int, k int) [][]int {
	combineSequence = make([][]int, 0)
	array := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		array = append(array, i)
	}
	combineExec(array, k, make([]int, 0, k))
	return combineSequence
}

func combineExec(array []int, k int, sequence []int) {
	if k == 0 {
		combineSequence = append(combineSequence, newSlice(sequence))
		return
	}
	for i := 0; i < len(array); i++ {
		combineExec(array[i+1:], k-1, append(sequence, array[i]))
	}
}

// 深拷贝
func newSlice(oldSlice []int) []int {
	slice := make([]int, len(oldSlice))
	copy(slice, oldSlice)
	return slice
}

/*
	题目链接: https://leetcode-cn.com/problems/combinations/submissions/
*/

/*
	总结
	1. 注意以上的代码求的是没有重复元素数组中的组合，有重复元素的话需要进行额外判断。
	2. 切片可以初始化容量，这样可以避免当切片长度大于容量时，切片扩容导致的额外代价。
	3. 该题题意是: 从数组中选取k个数的组合总数
*/