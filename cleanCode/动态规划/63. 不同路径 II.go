package 动态规划

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// 下面会遇到这样的操作: (obstacleGrid[x][y] ^ 1)
	// 当 obstacleGrid[x][y] == 0 时，(obstacleGrid[x][y] ^ 1) 返回 1
	// 当 obstacleGrid[x][y] == 1 时，(obstacleGrid[x][y] ^ 1) 返回 0
	// 通过这个运算，可以少写很多 if 判断

	if len(obstacleGrid) == 0 {
		return 0
	}
	rows, columns := len(obstacleGrid), len(obstacleGrid[0])
	countOfPath := get2DSlice(rows, columns)
	countOfPath[0][0] = obstacleGrid[0][0] ^ 1
	for row := 1; row < rows; row++ {
		countOfPath[row][0] = countOfPath[row-1][0] * (obstacleGrid[row][0] ^ 1)
	}
	for column := 1; column < columns; column++ {
		countOfPath[0][column] = countOfPath[0][column-1] * (obstacleGrid[0][column] ^ 1)
	}
	for row := 1; row < rows; row++ {
		for column := 1; column < columns; column++ {
			countOfPath[row][column] = (countOfPath[row-1][column] + countOfPath[row][column-1]) * (obstacleGrid[row][column] ^ 1)
		}
	}
	return countOfPath[rows-1][columns-1]
}

func get2DSlice(rows, column int) [][]int {
	slice := make([][]int, rows)
	for i := 0; i < len(slice); i++ {
		slice[i] = make([]int, column)
	}
	return slice
}


/*
	题目链接: https://leetcode-cn.com/problems/unique-paths-ii/
	总结
		1. 普通DP...
*/
