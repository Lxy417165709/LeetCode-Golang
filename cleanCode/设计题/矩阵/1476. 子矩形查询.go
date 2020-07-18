package 矩阵

// -------------------------------- SubrectangleQueries --------------------------------
// 执行用时：56 ms, 在所有 Go 提交中击败了92.31% 的用户
// 内存消耗：7.2 MB, 在所有 Go 提交中击败了100.00% 的用户
//
// 概述: 这是一种非暴力的解决方案，
// 		UpdateSubrectangle  时间复杂度: O(1)
// 		GetValue 			时间复杂度: O(x)，其中 x 为 UpdateSubrectangle 的调用次数。

type SubrectangleQueries struct {
	matrix           [][]int
	updateRectangles []*UpdateRectangle
}

func Constructor(matrix [][]int) SubrectangleQueries {
	return SubrectangleQueries{
		matrix: matrix,
		updateRectangles: make([]*UpdateRectangle, 0),
	}
}

func (sq *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	sq.updateRectangles = append(sq.updateRectangles, NewUpdateRectangle(row1, col1, row2, col2, newValue))
}

func (sq *SubrectangleQueries) GetValue(row int, col int) int {
	for i := len(sq.updateRectangles) - 1; i >= 0; i-- {
		if isCoordinateInUpdateRectangle(sq.updateRectangles[i], row, col) {
			return sq.updateRectangles[i].Value
		}
	}
	return sq.matrix[row][col]
}

func isCoordinateInUpdateRectangle(rectangle *UpdateRectangle, x, y int) bool {
	return x >= rectangle.LeftUpX && x <= rectangle.RightDownX && y >= rectangle.LeftUpY && y <= rectangle.RightDownY
}

// -------------------------------- UpdateRectangle --------------------------------
type UpdateRectangle struct {
	LeftUpX    int
	LeftUpY    int
	RightDownX int
	RightDownY int
	Value      int
}

func NewUpdateRectangle(row1 int, col1 int, row2 int, col2 int, value int) *UpdateRectangle {
	return &UpdateRectangle{row1, col1, row2, col2, value}
}

/*
	题目链接: https://leetcode-cn.com/problems/subrectangle-queries/
	总结:
		1. 在更新的时候，有时做一下标记就可以了，不用真实的进行更新。
*/
