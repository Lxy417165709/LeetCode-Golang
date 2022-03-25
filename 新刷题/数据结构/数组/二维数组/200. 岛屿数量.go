package 二维数组

import (
	"github.com/Lxy417165709/LeetCode-Golang/新刷题/util/matrix_util"
	"github.com/Lxy417165709/LeetCode-Golang/新刷题/util/model_util"
)

// numIslands 获取陆地数量。
func numIslands(grid [][]byte) int {
	countOfLand := 0
	for i := 0; i < len(grid); i++ {
		for t := 0; t < len(grid[i]); t++ {
			if grid[i][t] == '0' {
				continue
			}
			submergeLand(grid, i, t)
			countOfLand++
		}
	}
	return countOfLand
}

// submergeLand 淹没陆地。
func submergeLand(matrix [][]byte, x, y int) {
	// 1. 越界判断。
	height, width := matrix_util.GetHeightAndWidthOfByteMatrix(matrix)
	if x <= -1 || x >= height || y <= -1 || y >= width {
		return
	}

	// 2. 海判断。
	if matrix[x][y] == '0' {
		return
	}

	// 3. 淹没陆地，防止重复访问。
	matrix[x][y] = '0'

	// 4. DFS。
	offsets := []*model_util.Offset{
		{
			X: 1,
			Y: 0,
		},
		{
			X: -1,
			Y: 0,
		},
		{
			X: 0,
			Y: -1,
		},
		{
			X: 0,
			Y: 1,
		},
	}
	for _, offset := range offsets {
		submergeLand(matrix, x+offset.X, y+offset.Y)
	}
}
