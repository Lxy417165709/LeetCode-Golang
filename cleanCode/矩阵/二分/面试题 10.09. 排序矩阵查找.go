package 二分

// --------------------- 暴力查找法 ---------------------
// 略

// --------------------- 二分查找法 ---------------------
func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	for row := 0; row < rows; row++ {
		if isExist(matrix[row], target) {
			return true
		}
	}
	return false
}

func isExist(nums []int, ref int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == ref {
			return true
		}
		if nums[mid] > ref {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

// --------------------- 对角查找法 ---------------------
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	rows := len(matrix)
	cols := len(matrix[0])
	for row, col := 0, cols-1; row < rows && col >= 0; {
		if matrix[row][col] == target {
			return true
		}
		if matrix[row][col] > target {
			col--
		} else {
			row++
		}
	}
	return false
}

/*
	题目链接: https://leetcode-cn.com/problems/sorted-matrix-search-lcci/
*/
