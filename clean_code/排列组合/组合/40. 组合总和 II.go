package 组合

import "sort"

var combinations [][]int

func combinationSum2(candidates []int, target int) [][]int {
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

	haveCombinationsStartingWithThisNumFormed := make(map[int]bool)
	for i := 0; i < len(sortedArray); i++ {
		if haveCombinationsStartingWithThisNumFormed[sortedArray[i]] {
			continue
		}
		haveCombinationsStartingWithThisNumFormed[sortedArray[i]] = true
		formCombinations(sortedArray[i+1:], append(nowCombinations, sortedArray[i]), nowSum+sortedArray[i], sumOfWantToSelect)
	}
}

func NewSlice(oldSlice []int) []int {
	newSlice := make([]int, len(oldSlice))
	copy(newSlice, oldSlice)
	return newSlice
}

/*
	题目链接: https://leetcode-cn.com/problems/combination-sum-ii/
	总结:
		1. 求取数组组合之前，需要保证数组有序。
		2. 理解的还不够透彻 2020年7月8日14:34:24
*/
