package main

func maxIncreaseKeepingSkyline(grid [][]int) int {
	rows, cols := getRowsAndCols(grid)
	highOfSkyLineInRow := make([]int, rows)
	highOfSkyLineInCol := make([]int, cols)

	for i := 0; i < rows; i++ {
		for t := 0; t < cols; t++ {
			highOfSkyLineInRow[i] = max(highOfSkyLineInRow[i], grid[i][t])
			highOfSkyLineInCol[t] = max(highOfSkyLineInCol[t], grid[i][t])
		}
	}
	maxIncrease := 0
	for i := 0; i < rows; i++ {
		for t := 0; t < cols; t++ {
			maxIncrease += min(highOfSkyLineInCol[t], highOfSkyLineInRow[i]) - grid[i][t]
		}
	}
	return maxIncrease

}


func getRowsAndCols(matrix [][]int) (int, int) {
	if len(matrix) == 0 {
		return 0, 0
	}
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
	题目链接: https://leetcode-cn.com/problems/max-increase-to-keep-city-skyline/
*/
