
package 组合

import "sort"

func subsetsWithDup(nums []int) [][]int {
	combinations = make([][]int, 0)
	sort.Ints(nums)
	formCombinations(nums,[]int{})
	return combinations
}


var combinations [][]int
func formCombinations(sortedArray []int, nowCombinations []int) {
	combinations = append(combinations, NewSlice(nowCombinations))
	haveCombinationsStartingWithThisNumFormed:= make(map[int]bool)
	for i := 0; i < len(sortedArray); i++ {
		if haveCombinationsStartingWithThisNumFormed[sortedArray[i]]{
			continue
		}
		haveCombinationsStartingWithThisNumFormed[sortedArray[i]] = true
		formCombinations(sortedArray[i+1:], append(nowCombinations, sortedArray[i]))
	}
}

func NewSlice(oldSlice []int) []int {
	newSlice := make([]int, len(oldSlice))
	copy(newSlice, oldSlice)
	return newSlice
}

/*
	题目链接: https://leetcode-cn.com/problems/subsets/submissions/
*/
