package 组合

import "sort"

var combinations [][]int

func combinationSum(candidates []int, target int) [][]int {
	combinations = make([][]int, 0)
	sort.Ints(candidates)
	formCombinations(candidates, []int{}, 0, target)
	return combinations
}


func formCombinations(sortedArray []int, nowCombinations []int, nowSum int, sumOfWantToSelect int) {
	if nowSum > sumOfWantToSelect {
		return // 题目说了: 所有数字（包括目标数）都是正整数。
	}
	if nowSum == sumOfWantToSelect {
		combinations = append(combinations, NewSlice(nowCombinations))
		return
	}
	for i := 0; i < len(sortedArray); i++ {
		// 组合总和II：sortedArray[i+1:]
		// 组合总和：sortedArray[i:]
		formCombinations(sortedArray[i:], append(nowCombinations, sortedArray[i]), nowSum+sortedArray[i], sumOfWantToSelect)
	}
}

func NewSlice(oldSlice []int) []int {
	newSlice := make([]int, len(oldSlice))
	copy(newSlice, oldSlice)
	return newSlice
}

/*
	题目链接: https://leetcode-cn.com/problems/combination-sum/
	总结
		1. 注意与组合总和II相区别，本题数字可以选无数次、而且数组没有重复数字。
*/
