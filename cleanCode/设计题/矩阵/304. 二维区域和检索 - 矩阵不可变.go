package 矩阵
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

func (nm *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return nm.getAreaSum(row1, col1, row2, col2)
}

func (nm *NumMatrix) getAreaSum(row1 int, col1 int, row2 int, col2 int) int {
	rows, cols := getRowsAndCols(nm.suffixSumMatrix)
	if rows == 0 {
		return 0
	}
	sum1, sum2, sum3, sum4 := 0, 0, 0, 0
	switch {
	case row1-1 >= 0 && col1-1 >= 0:
		sum1 = nm.suffixSumMatrix[row1-1][col1-1]
		sum2 = nm.suffixSumMatrix[row1-1][min(cols, col2)]
		sum3 = nm.suffixSumMatrix[min(rows, row2)][col1-1]
	case col1-1 >= 0:
		sum3 = nm.suffixSumMatrix[min(rows, row2)][col1-1]
	case row1-1 >= 0:
		sum2 = nm.suffixSumMatrix[row1-1][min(cols, col2)]
	}
	sum4 = nm.suffixSumMatrix[min(rows, row2)][min(cols, col2)]
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
	题目链接: https://leetcode-cn.com/problems/range-sum-query-2d-immutable/submissions/
*/
