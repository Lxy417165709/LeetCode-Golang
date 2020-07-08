package 组合

import "sort"

var combinations [][]int

func combinationSum3(k int, n int) [][]int {
	combinations = make([][]int, 0)
	candidates := getCandidates()
	formCombinations(candidates, []int{}, 0, k, n)
	return combinations
}

func getCandidates() []int {
	candidates := make([]int, 0)
	for i := 1; i <= 9; i++ {
		candidates = append(candidates, i)
	}
	return candidates
}

func formCombinations(sortedArray []int, nowCombinations []int, nowSum int, countOfWantToSelect, sumOfWantToSelect int) {
	if nowSum > sumOfWantToSelect || len(nowCombinations) > countOfWantToSelect {
		return
	}
	if nowSum == sumOfWantToSelect {
		if countOfWantToSelect == len(nowCombinations) {
			combinations = append(combinations, NewSlice(nowCombinations))
		}
		return
	}
	for i := 0; i < len(sortedArray); i++ {
		formCombinations(
			sortedArray[i+1:],
			append(nowCombinations, sortedArray[i]),
			nowSum+sortedArray[i],
			countOfWantToSelect,
			sumOfWantToSelect,
		)
	}
}

func NewSlice(oldSlice []int) []int {
	newSlice := make([]int, len(oldSlice))
	copy(newSlice, oldSlice)
	return newSlice
}

/*
	题目链接: https://leetcode-cn.com/problems/combination-sum-iii/
*/
