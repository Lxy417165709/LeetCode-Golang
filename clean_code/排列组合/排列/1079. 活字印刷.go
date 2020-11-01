package main

func numTilePossibilities(tiles string) int {
	permutations = make([][]int, 0)
	formPermutations(stringToIntSlice(tiles), []int{})
	return len(permutations) - 1 // -1是因为有个空的
}

var permutations [][]int

func formPermutations(array []int, nowPermutation []int) {
	permutations = append(permutations, NewSlice(nowPermutation))
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

func stringToIntSlice(s string) []int {
	slice := make([]int, 0)
	for i := 0; i < len(s); i++ {
		slice = append(slice, int(s[i]))
	}
	return slice
}

func NewSlice(oldSlice []int) []int {
	newSlice := make([]int, len(oldSlice))
	copy(newSlice, oldSlice)
	return newSlice
}

/*
	题目链接: https://leetcode-cn.com/problems/letter-tile-possibilities/
	总结
	1. 这题其实不用得到所有具体的全排列，得到所有全排列的个数就可以了，可以从这方面优化。
	2. 官方还有更好的题解。
*/
