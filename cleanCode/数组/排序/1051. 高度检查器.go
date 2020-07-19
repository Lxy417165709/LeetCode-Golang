package 排序

import "sort"

func heightChecker(heights []int) int {
	sortedHeights := NewSlice(heights)
	sort.Ints(sortedHeights)
	countOfMustMovePerson := 0
	for i := 0; i < len(heights); i++ {
		if heights[i] != sortedHeights[i] {
			countOfMustMovePerson++
		}
	}
	return countOfMustMovePerson
}

func NewSlice(oldSlice []int) []int {
	newSlice := make([]int, len(oldSlice))
	copy(newSlice, oldSlice)
	return newSlice
}


/*
	题目链接: https://leetcode-cn.com/problems/height-checker/
*/
