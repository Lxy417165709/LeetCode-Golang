package 动态规划

import "fmt"

// --------------------- 未封装为对象 ---------------------
func matrixBlockSum(mat [][]int, K int) [][]int {
	suffixSumMatrix := getSuffixSumMatrix(mat)
	return getAreaSumMatrix(mat, suffixSumMatrix, K)
}

func getAreaSumMatrix(mat [][]int, suffixSumMatrix [][]int, K int) [][]int {
	rows, cols := getRowsAndCols(mat)
	areaSumMatrix := get2DSlice(rows, cols)
	for i := 0; i < rows; i++ {
		for t := 0; t < cols; t++ {
			areaSumMatrix[i][t] = getAreaSum(mat, suffixSumMatrix, i, t, K)
		}
	}
	return areaSumMatrix
}

func getSuffixSumMatrix(mat [][]int) [][]int {
	rows, cols := getRowsAndCols(mat)
	suffixSumMatrix := get2DSlice(rows, cols)
	for i := 0; i < rows; i++ {
		for t := 0; t < cols; t++ {
			suffixSumMatrix[i][t] = getSuffixSum(mat, suffixSumMatrix, i, t)
		}
	}
	return suffixSumMatrix
}

func getAreaSum(mat [][]int, suffixSumMatrix [][]int, x, y, K int) int {
	rows, cols := getRowsAndCols(mat)
	sum1, sum2, sum3, sum4 := 0, 0, 0, 0
	switch {
	case x-K-1 >= 0 && y-K-1 >= 0:
		sum1 = suffixSumMatrix[x-K-1][y-K-1]
		sum2 = suffixSumMatrix[x-K-1][min(y+K, cols-1)]
		sum3 = suffixSumMatrix[min(x+K, rows-1)][y-K-1]
	case x-K-1 >= 0:
		sum2 = suffixSumMatrix[x-K-1][min(y+K, cols-1)]
	case y-K-1 >= 0:
		sum3 = suffixSumMatrix[min(x+K, rows-1)][y-K-1]
	}
	sum4 = suffixSumMatrix[min(x+K, rows-1)][min(y+K, cols-1)]
	return sum4 - sum2 - sum3 + sum1
}

func getSuffixSum(mat [][]int, suffixSumMatrix [][]int, x, y int) int {
	sum1, sum2, sum3 := 0, 0, 0
	switch {
	case x >= 1 && y >= 1:
		sum3 = suffixSumMatrix[x-1][y-1]
		sum2 = suffixSumMatrix[x-1][y]
		sum1 = suffixSumMatrix[x][y-1]
	case x >= 1:
		sum2 = suffixSumMatrix[x-1][y]
	case y >= 1:
		sum1 = suffixSumMatrix[x][y-1]
	}
	return mat[x][y] + sum1 + sum2 - sum3
}

func getRowsAndCols(mat [][]int) (int, int) {
	return len(mat), len(mat[0])
}

func get2DSlice(rows, column int) [][]int {
	slice := make([][]int, rows)
	for i := 0; i < len(slice); i++ {
		slice[i] = make([]int, column)
	}
	return slice
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// --------------------- 封装为对象 ---------------------
// --------------------- 封装为对象 ---------------------
func matrixBlockSum(mat [][]int, K int) [][]int {
	numMatrix := Constructor(mat)
	rows, cols := getRowsAndCols(mat)
	areaSumMatrix := get2DSlice(rows, cols)
	for i := 0; i < rows; i++ {
		for t := 0; t < cols; t++ {
			areaSumMatrix[i][t] = numMatrix.GetAreaSum(i-K, t-K, i+K, t+K)
		}
	}
	return areaSumMatrix
}

// -------------------- NumMatrix --------------------
type NumMatrix struct {
	matrix          [][]int
	suffixSumMatrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	numMatrix := NumMatrix{
		matrix: matrix,
	}
	numMatrix.formSuffixSumMatrix()
	return numMatrix
}
func (nm *NumMatrix) formSuffixSumMatrix() {
	rows, cols := getRowsAndCols(nm.matrix)
	nm.suffixSumMatrix = get2DSlice(rows, cols)
	for i := 0; i < rows; i++ {
		for t := 0; t < cols; t++ {
			nm.suffixSumMatrix[i][t] = nm.getSuffixSum(i, t)
		}
	}
}

func (nm *NumMatrix) GetAreaSum(row1 int, col1 int, row2 int, col2 int) int {
	rows, cols := getRowsAndCols(nm.suffixSumMatrix)
	if rows == 0 {
		return 0
	}
	sum1, sum2, sum3, sum4 := 0, 0, 0, 0
	switch {
	case row1-1 >= 0 && col1-1 >= 0:
		sum1 = nm.suffixSumMatrix[row1-1][col1-1]
		sum2 = nm.suffixSumMatrix[row1-1][min(cols-1, col2)]
		sum3 = nm.suffixSumMatrix[min(rows-1, row2)][col1-1]
	case col1-1 >= 0:
		sum3 = nm.suffixSumMatrix[min(rows-1, row2)][col1-1]
	case row1-1 >= 0:
		sum2 = nm.suffixSumMatrix[row1-1][min(cols-1, col2)]
	}
	sum4 = nm.suffixSumMatrix[min(rows-1, row2)][min(cols-1, col2)]
	return sum4 - sum2 - sum3 + sum1
}

func (nm *NumMatrix) getSuffixSum(x, y int) int {
	sum1, sum2, sum3 := 0, 0, 0
	switch {
	case x >= 1 && y >= 1:
		sum3 = nm.suffixSumMatrix[x-1][y-1]
		sum2 = nm.suffixSumMatrix[x-1][y]
		sum1 = nm.suffixSumMatrix[x][y-1]
	case x >= 1:
		sum2 = nm.suffixSumMatrix[x-1][y]
	case y >= 1:
		sum1 = nm.suffixSumMatrix[x][y-1]
	}
	return nm.matrix[x][y] + sum1 + sum2 - sum3
}

func getRowsAndCols(mat [][]int) (int, int) {
	if len(mat) == 0 {
		return 0, 0
	}
	return len(mat), len(mat[0])
}

func get2DSlice(rows, column int) [][]int {
	slice := make([][]int, rows)
	for i := 0; i < len(slice); i++ {
		slice[i] = make([]int, column)
	}
	return slice
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*
	题目链接: https://leetcode-cn.com/problems/matrix-block-sum/comments/
	总结:
		1. 这题就是求矩阵的前缀和。
		2. 要进行分析，不然一点错误都要找好久~
*/
