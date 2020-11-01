package 数组

import "sort"

// ------------------------ 排序解法 ------------------------
// 时间复杂度: O(n * log_n)
func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	return isArithmeticProgression(arr)
}

func isArithmeticProgression(arr []int) bool {
	if len(arr) <= 1 {
		return true
	}
	diff := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != diff {
			return false
		}
	}
	return true
}

/*
	题目链接: https://leetcode-cn.com/problems/can-make-arithmetic-progression-from-sequence/
*/
