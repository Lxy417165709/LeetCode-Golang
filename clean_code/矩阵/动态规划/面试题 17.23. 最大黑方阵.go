package 动态规划

func findSquare(matrix [][]int) []int {
	x, y, size := getMaxBlankSquareCoordinateAndSize(matrix)
	if x == -1 {
		return []int{}
	}
	return []int{x - size + 1, y - size + 1, size}
}

func getMaxBlankSquareCoordinateAndSize(matrix [][]int) (int, int, int) {
	maxWidth, maxHeight := getMaxWidthAndMaxHeight(matrix)
	rows, cols := getRowsAndCols(matrix)
	maxBlankSquareX, maxBlankSquareY, maxSize := -1, -1, 0
	for i := 0; i < rows; i++ {
		for t := 0; t < cols; t++ {
			curBlankSquareSize := getBlankSquareSize(maxWidth, maxHeight, i, t)
			if curBlankSquareSize > maxSize {
				maxBlankSquareX = i
				maxBlankSquareY = t
				maxSize = curBlankSquareSize
			}
		}
	}
	return maxBlankSquareX, maxBlankSquareY, maxSize
}

func getMaxWidthAndMaxHeight(matrix [][]int) ([][]int, [][]int) {
	rows, cols := getRowsAndCols(matrix)
	maxWidth := get2DSlice(rows, cols)
	maxHeight := get2DSlice(rows, cols)
	maxWidth[0][0] = matrix[0][0] ^ 1
	maxHeight[0][0] = matrix[0][0] ^ 1
	for i := 1; i < rows; i++ {
		maxWidth[i][0] = matrix[i][0] ^ 1
		if matrix[i][0] == 0 {
			maxHeight[i][0] = maxHeight[i-1][0] + 1
		}
	}
	for t := 1; t < cols; t++ {
		maxHeight[0][t] = matrix[0][t] ^ 1
		if matrix[0][t] == 0 {
			maxWidth[0][t] = maxWidth[0][t-1] + 1
		}
	}
	for i := 1; i < rows; i++ {
		for t := 1; t < cols; t++ {
			if matrix[i][t] == 0 {
				maxHeight[i][t] = maxHeight[i-1][t] + 1
				maxWidth[i][t] = maxWidth[i][t-1] + 1
			}
		}
	}
	return maxWidth, maxHeight
}

func getBlankSquareSize(maxWidth, maxHeight [][]int, x, y int) int {
	maxSize := 0
	for size := 1; size <= min(maxHeight[x][y], maxWidth[x][y]); size++ {
		if maxHeight[x][y-size+1] >= size && maxWidth[x-size+1][y] >= size {
			maxSize = size
		}
	}
	return maxSize
}

func getRowsAndCols(matrix [][]int) (int, int) {
	if len(matrix) == 0 {
		return 0, 0
	}
	return len(matrix), len(matrix[0])
}

func get2DSlice(rows, column int) [][]int {
	slice := make([][]int, rows)
	for i := 0; i < len(slice); i++ {
		slice[i] = make([]int, column)
	}
	return slice
}

func min(arr ...int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	a, b := arr[0], min(arr[1:]...)
	if a > b {
		return b
	}
	return a
}

/*
	题目链接: https://leetcode-cn.com/problems/max-black-square-lcci/
	总结;
		1. 这题要找的是 4 条边为 0 的方阵，该方阵可以不是实心的。 而 _1277. 统计全为1的正方形子矩阵_ 它
			要求矩阵是实心的。
*/
