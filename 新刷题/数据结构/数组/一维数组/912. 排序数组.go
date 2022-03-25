package 一维数组

import "github.com/Lxy417165709/LeetCode-Golang/新刷题/util/sort_util"

// sortArray 快速排序。
func sortArray(nums []int) []int {
	sort_util.QuickSort(nums)
	return nums
}

// sortArrayByHeap 堆排序。
func sortArrayByHeap(nums []int) []int {
	return sort_util.HeapSort(nums)
}
