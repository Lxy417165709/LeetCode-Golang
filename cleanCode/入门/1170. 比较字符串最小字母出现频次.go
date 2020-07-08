package 入门

import (
	"sort"
)

func numSmallerByFrequency(queries []string, words []string) []int {
	queryFrequencyArray := getFrequencyArray(queries)
	wordsFrequencyArray := getFrequencyArray(words)
	sort.Ints(wordsFrequencyArray)
	result := make([]int, 0)
	for i := 0; i < len(queryFrequencyArray); i++ {
		result = append(result, getCountOfGreaterElement(wordsFrequencyArray, queryFrequencyArray[i]))
	}
	return result
}

func getCountOfGreaterElement(nums []int, ref int) int {
	return len(nums) - getIndexOfFirstGreater(nums, ref)
}

func getIndexOfFirstGreater(nums []int, ref int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > ref {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func getFrequencyArray(words []string) []int {
	array := make([]int, 0)
	for _, word := range words {
		array = append(array, getFrequencyOfMinChar(word))
	}
	return array
}

func getFrequencyOfMinChar(word string) int {
	countOfChar := make(map[byte]int)
	for i := 0; i < len(word); i++ {
		countOfChar[word[i]]++
	}
	var minChar byte = 'z' + 1
	frequency := 0
	for char, count := range countOfChar {
		if minChar > char {
			frequency = count
			minChar = char
		}
	}
	return frequency
}


/*
	题目链接: https://leetcode-cn.com/problems/compare-strings-by-frequency-of-the-smallest-character/
	总结
		1. 这题题意有点难懂，操作还是很简单的
*/
