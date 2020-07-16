package 动态规划

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

/*
	题目链接: https://leetcode-cn.com/problems/matrix-block-sum/comments/
	总结:
		1. 这题就是求矩阵的前缀和。
*/
