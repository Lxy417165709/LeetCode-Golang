package 二维数组

import (
	"github.com/Lxy417165709/LeetCode-Golang/新刷题/matrix_util"
)

func findNumberIn2DArray(matrix [][]int, target int) bool {
	height, width := matrix_util.GetHeightAndWidth(matrix)
	column, row := width-1, 0
	for column >= 0 && row <= height-1 {
		reference := matrix[row][column]
		if reference == target {
			return true
		}
		if reference > target {
			column--
		}
		if reference < target {
			row++
		}
	}
	return false
}
