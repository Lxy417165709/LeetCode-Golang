package 一维数组

import "github.com/Lxy417165709/LeetCode-Golang/新刷题/util/sort_util"

func sortArray(nums []int) []int {
	sort_util.QuickSort(nums)
	return nums
}


