package main

func uniquePathsIII(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	beginX, beginY := -1, -1
	validGrid, validPath = 0, 0
	dx = []int{0, 0, -1, 1}
	dy = []int{1, -1, 0, 0}
	for i := 0; i < len(grid); i++ {
		for t := 0; t < len(grid[i]); t++ {
			if grid[i][t] == -1 {
				continue
			}
			if grid[i][t] == 1 {
				beginX, beginY = i, t
			}
			validGrid++
		}
	}
	uniquePathsIIIExec(grid, beginX, beginY)
	return validPath

}

var validGrid int // 0,1,2总的格子数
var validPath int // 合法路径数
var dx, dy []int  // 移动方向向量

// DFS
func uniquePathsIIIExec(grid [][]int, x, y int) {
	if grid[x][y] == 2 {
		// 这个 validGrid 就是值为2的格子
		if validGrid == 1 {
			validPath++
		}
		return
	}
	grid[x][y] = -1
	validGrid--
	for i := 0; i < len(dx); i++ {
		nx, ny := x+dx[i], y+dy[i]
		if nx < 0 || ny < 0 || nx >= len(grid) || ny >= len(grid[nx]) || grid[nx][ny] == -1 {
			continue
		}
		uniquePathsIIIExec(grid, nx, ny)
	}
	validGrid++
	grid[x][y] = 0
}

/*
	总结
	1. 第一次看到这个题的时，我其实不会做...我是看了官方题解，发现是DFS后我才做出来的。
	2. 以上解法是采用深度优先搜索策略的回溯解法。
	3. 以上可以当做一个DFS模板。
*/
