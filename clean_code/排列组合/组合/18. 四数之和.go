package 组合

import "sort"

//	------------------------- 注意以下代码不会错误，但是会超时 --------------------
func fourSum(nums []int, target int) [][]int {
	combinations = make([][]int, 0)
	sort.Ints(nums)
	formCombinations(nums, []int{}, 0, 4, target)
	return combinations
}

var combinations [][]int

func formCombinations(sortedArray []int, nowCombinations []int, nowSum int, countOfWantToSelect, sumOfWantToSelect int) {
	if len(nowCombinations) > countOfWantToSelect {
		return
	}
	if nowSum == sumOfWantToSelect && countOfWantToSelect == len(nowCombinations) {
		combinations = append(combinations, NewSlice(nowCombinations))
		return
	}

	haveCombinationsStartingWithThisNumFormed := make(map[int]bool)
	for i := 0; i < len(sortedArray); i++ {
		if haveCombinationsStartingWithThisNumFormed[sortedArray[i]] {
			continue
		}
		haveCombinationsStartingWithThisNumFormed[sortedArray[i]] = true
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
