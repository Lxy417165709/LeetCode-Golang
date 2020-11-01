package main

func spiralMatrixIII(R int, C int, r0 int, c0 int) [][]int {
	leftBound, rightBound := c0, c0
	upBound, downBound := r0, r0
	result := make([][]int, 0)
	for len(result) < R*C {
		for i := leftBound; i <= rightBound; i++ {
			if isValidCoordinate(upBound, i, R, C) {
				result = append(result, []int{upBound, i})
			}
		}
		rightBound++
		for i := upBound; i <= downBound; i++ {
			if isValidCoordinate(i, rightBound, R, C) {
				result = append(result, []int{i, rightBound})
			}
		}
		downBound++
		for i := rightBound; i >= leftBound; i-- {
			if isValidCoordinate(downBound, i, R, C) {
				result = append(result, []int{downBound, i})
			}
		}
		leftBound--
		for i := downBound; i >= upBound; i-- {
			if isValidCoordinate(i, leftBound, R, C) {
				result = append(result, []int{i, leftBound})
			}
		}
		upBound--
	}
	return result
}

func isValidCoordinate(x, y, R, C int) bool {
	return x >= 0 && x <= R-1 && y >= 0 && y <= C-1
}


/*
	题目链接: https://leetcode-cn.com/problems/spiral-matrix-iii/
	总结:
		1. 其实关于矩阵的题，可以封装为一个结构，这样就能减少函数内的代码，也能减少函数参数个数。
*/
