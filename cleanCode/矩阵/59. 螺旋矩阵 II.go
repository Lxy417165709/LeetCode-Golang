package main

func generateMatrix(n int) [][]int {
	matrix := get2DSlice(n, n)
	leftBound, rightBound := 0, n-1
	upBound, downBound := 0, n-1
	curNumber := 0
	for leftBound <= rightBound && upBound <= downBound {
		for i := leftBound; i <= rightBound; i++ {
			curNumber++
			matrix[upBound][i] = curNumber
		}
		upBound++
		for i := upBound; i <= downBound; i++ {
			curNumber++
			matrix[i][rightBound] = curNumber
		}
		rightBound--
		for i := rightBound; i >= leftBound && upBound <= downBound; i-- {
			curNumber++
			matrix[downBound][i] = curNumber
		}
		downBound--
		for i := downBound; i >= upBound && leftBound <= rightBound; i-- {
			curNumber++
			matrix[i][leftBound] = curNumber
		}
		leftBound++
	}
	return matrix
}

func get2DSlice(rows, column int) [][]int {
	slice := make([][]int, rows)
	for i := 0; i < len(slice); i++ {
		slice[i] = make([]int, column)
	}
	return slice
}

/*
	题目链接; https://leetcode-cn.com/problems/spiral-matrix-ii/
*/
