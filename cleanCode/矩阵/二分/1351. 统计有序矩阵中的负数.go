package 二分

func countNegatives(grid [][]int) int {
	countOfNegativeNum := 0
	for i := 0; i < len(grid); i++ {
		countOfNegativeNum += len(grid[i]) - getIndexOfLastLess(grid[i], 0)
	}
	return countOfNegativeNum
}

func getIndexOfLastLess(array []int, ref int) int {
	left, right := 0, len(array)-1
	for left <= right {
		mid := left + (right-left)/2
		if array[mid] == ref {
			left = mid + 1
		} else {
			if array[mid] > ref {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return left
}

/*
	题目链接: https://leetcode-cn.com/problems/count-negative-numbers-in-a-sorted-matrix/
	总结:
		1. 这题和以往不一样的地方就是: 矩阵每一行是递减的，而不是递增的，所以二分函数也要进行相应修改
*/
