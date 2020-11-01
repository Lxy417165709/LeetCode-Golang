package main

const INF = 10000000000000

func luckyNumbers(matrix [][]int) []int {
	minValueInRow, maxValueInCol := getMinValueInRowAndMaxValueInCol(matrix)
	return getLuckyNumbers(matrix, minValueInRow, maxValueInCol)
}

func getMinValueInRowAndMaxValueInCol(matrix [][]int) ([]int, []int) {
	rows, cols := getRowsAndCols(matrix)
	minValueInRow, maxValueInCol := make([]int, rows), make([]int, cols)
	for i := 0; i < rows; i++ {
		minValueInRow[i] = INF
	}
	for t := 0; t < cols; t++ {
		maxValueInCol[t] = -INF
	}
	for i := 0; i < rows; i++ {
		for t := 0; t < cols; t++ {
			minValueInRow[i] = min(minValueInRow[i], matrix[i][t])
			maxValueInCol[t] = max(maxValueInCol[t], matrix[i][t])
		}
	}
	return minValueInRow, maxValueInCol
}

func getLuckyNumbers(matrix [][]int, minValueInRow, maxValueInCol []int) []int {
	rows, cols := getRowsAndCols(matrix)
	luckyNums := make([]int, 0)
	for i := 0; i < rows; i++ {
		for t := 0; t < cols; t++ {
			if matrix[i][t] == minValueInRow[i] && matrix[i][t] == maxValueInCol[t] {
				luckyNums = append(luckyNums, matrix[i][t])
			}
		}
	}
	return luckyNums
}

func getRowsAndCols(matrix [][]int) (int, int) {
	return len(matrix), len(matrix[0])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*
	题目链接: https://leetcode-cn.com/problems/lucky-numbers-in-a-matrix/
*/
