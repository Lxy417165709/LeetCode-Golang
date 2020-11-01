package 贪心

func matrixScore(matrix [][]int) int {
	rows, cols := getRowsAndCols(matrix)
	for i := 0; i < rows; i++ {
		if isFirstElementZero(matrix, i) {
			moveRow(matrix, i)
		}
	}
	for t := 0; t < cols; t++ {
		if isZeroMoreThanOne(matrix, t) {
			moveCol(matrix, t)
		}
	}
	return getBinaryFromMatrix(matrix)
}

func isFirstElementZero(matrix [][]int, row int) bool {
	return matrix[row][0] == 0
}

func moveRow(matrix [][]int, row int) {
	_, cols := getRowsAndCols(matrix)
	for i := 0; i < cols; i++ {
		matrix[row][i] ^= 1
	}
}

func isZeroMoreThanOne(matrix [][]int, col int) bool {
	rows, _ := getRowsAndCols(matrix)
	countOfZero := 0
	for i := 0; i < rows; i++ {
		if matrix[i][col] == 0 {
			countOfZero++
		}
	}
	return countOfZero > rows-countOfZero
}

func moveCol(matrix [][]int, col int) {
	rows, _ := getRowsAndCols(matrix)
	for i := 0; i < rows; i++ {
		matrix[i][col] ^= 1
	}
}

func getBinaryFromMatrix(matrix [][]int) int {
	rows, _ := getRowsAndCols(matrix)
	sum := 0
	for i := 0; i < rows; i++ {
		sum += getBinaryFromArray(matrix[i])
	}
	return sum
}

func getBinaryFromArray(array []int) int {
	sum := 0
	for i := 0; i < len(array); i++ {
		sum = sum*2 + array[i]
	}
	return sum
}

func getRowsAndCols(matrix [][]int) (int, int) {
	return len(matrix), len(matrix[0])
}

/*
	题目链接: https://leetcode-cn.com/problems/score-after-flipping-matrix/
	总结:
		1. 二进制最高位的数字一定大于低于该位的所有位的和。
*/
