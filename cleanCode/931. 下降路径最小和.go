package main

func minFallingPathSum(A [][]int) int {
	rows,columns := len(A),len(A[0])
	minSum := get2DSlice(rows, columns)
	for column := 0; column < columns; column++ {
		minSum[0][column] = A[0][column]
	}
	for row := 1; row < rows; row++ {
		for column := 0; column < columns; column++ {
			minSum[row][column] = min(minSum[row-1][column])
			if column-1 >= 0 {
				minSum[row][column] = min(minSum[row][column], minSum[row-1][column-1])
			}
			if column+1 <= columns-1 {
				minSum[row][column] = min(minSum[row][column], minSum[row-1][column+1])
			}
			minSum[row][column] += A[row][column]
		}
	}
	return min(minSum[rows-1]...)
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
	a := arr[0]
	b := min(arr[1:]...)
	if a > b {
		return b
	}
	return a
}
