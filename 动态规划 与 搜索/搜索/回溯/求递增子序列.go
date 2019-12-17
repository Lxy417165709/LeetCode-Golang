package main

/*
	给定一个整型数组, 你的任务是找到所有该数组的递增子序列，递增子序列的长度至少是2。
*/

// 回溯 + 外部变量 获取数组中的所有长度大于1的递增子序列
var subsequences [][]int
func findSubsequences(nums []int) [][]int {
	subsequences = make([][]int, 0)
	findSubsequencesExec(nums, []int{})
	return subsequences
}

func findSubsequencesExec(nums []int, sequence []int) {
	if len(sequence) >= 2 {
		subsequences = append(subsequences, newSlice(sequence))
	}
	isVisit := make(map[int]bool)	// 为了去重，比如: 加入isVisit[5]==true,那么5就不能再被访问了。
	for i := 0; i < len(nums); i++ {
		if isVisit[nums[i]] == true {
			continue
		}
		isVisit[nums[i]] = true
		if len(sequence) == 0 || sequence[len(sequence)-1] <= nums[i] {
			findSubsequencesExec(nums[i+1:], append(sequence, nums[i]))
		}
	}
}

// 深复制
func newSlice(slice []int) []int {
	s := make([]int, len(slice))
	copy(s, slice)
	return s
}

/*
	题目链接:
		https://leetcode-cn.com/problems/increasing-subsequences/		递增子序列
*/
