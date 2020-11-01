package 双指针

import "sort"

// ----------------------------- 自己写的 -----------------------------
func findUnsortedSubarray(nums []int) int {
	prefixMax := getPrefixMax(nums)
	suffixMin := getSuffixMin(nums)
	left, right := 0, len(nums)-1
	for left <= right && nums[left] <= suffixMin[left] {
		left++
	}
	for left <= right && nums[right] >= prefixMax[right] {
		right--
	}
	return right - left + 1
}

func getPrefixMax(arr []int) []int {
	prefixMax := make([]int, len(arr))
	prefixMax[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		prefixMax[i] = max(prefixMax[i-1], arr[i])
	}
	return prefixMax
}

func getSuffixMin(arr []int) []int {
	suffixMin := make([]int, len(arr))
	suffixMin[len(arr)-1] = arr[len(arr)-1]
	for i := len(arr) - 2; i >= 0; i-- {
		suffixMin[i] = min(suffixMin[i+1], arr[i])
	}
	return suffixMin
}

func max(arr ...int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	a := arr[0]
	b := max(arr[1:]...)
	if a < b {
		return b
	}
	return a
}
func min(arr ...int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	a := arr[0]
	b := min(arr[1:]...)
	if a > b {
		return b
	}
	return a
}

// ----------------------------- 看了官方题解后写的 -------------------------------
func findUnsortedSubarray(nums []int) int {
	sortedNums := getSortedNums(nums)
	left, right := 0, len(nums)-1
	for left <= right && nums[left] <= sortedNums[left] {
		left++
	}
	for left <= right && nums[right] >= sortedNums[right] {
		right--
	}
	return right - left + 1
}

func getSortedNums(nums []int) []int {
	sortNums := NewSlice(nums)
	sort.Ints(sortNums)
	return sortNums
}

func NewSlice(oldSlice []int) []int {
	newSlice := make([]int, len(oldSlice))
	copy(newSlice, oldSlice)
	return newSlice
}

/*
	题目链接: https://leetcode-cn.com/problems/shortest-unsorted-continuous-subarray/
	总结:
		1. 这题运用了前缀最小、前缀最大 + 双指针
		2. 官方还有更优美的解法，运用排序。
		3. 官方还有O(1)空间的解法。
*/
