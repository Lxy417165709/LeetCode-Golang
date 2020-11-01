package main

var permutations [][]int

func permuteUnique(nums []int) [][]int {
	permutations = make([][]int, 0)
	formPermutations(nums, []int{})
	return permutations
}

func formPermutations(array []int, nowPermutation []int) {
	if len(array) == 0 {
		permutations = append(permutations, NewSlice(nowPermutation))
		return
	}
	havePermutationsStartingWithThisNumFormed := make(map[int]bool)
	for i := 0; i < len(array); i++ {
		if havePermutationsStartingWithThisNumFormed[array[i]] == true {
			continue
		}
		havePermutationsStartingWithThisNumFormed[array[i]] = true
		array[0], array[i] = array[i], array[0]
		formPermutations(array[1:], append(nowPermutation, array[0]))
		array[0], array[i] = array[i], array[0]
	}
}

func NewSlice(oldSlice []int) []int {
	newSlice := make([]int, len(oldSlice))
	copy(newSlice, oldSlice)
	return newSlice
}


/*
	题目链接: https://leetcode-cn.com/problems/permutations-ii/
	总结
		1. 题意: 给定一个可包含重复数字的序列，返回所有不重复的全排列。
*/
