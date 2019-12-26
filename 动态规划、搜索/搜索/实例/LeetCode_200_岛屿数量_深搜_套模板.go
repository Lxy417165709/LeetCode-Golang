package main

var dx []int
var dy []int

// DFS 解决矩阵连通块问题 (原题是求岛屿数量)
func numIslands(grid [][]byte) int {

	// 确定搜索方向
	dx = []int{0, 0, 1, -1}
	dy = []int{1, -1, 0, 0}

	ans := 0
	for i := 0; i < len(grid); i++ {
		for t := 0; t < len(grid[i]); t++ {
			if i < 0 || t < 0 || i >= len(grid) || t >= len(grid[i]) || grid[i][t] == '0' {
				continue
			}
			// 遇到陆地则进行DFS拓展
			ans++
			DFS(grid, i, t)
		}
	}
	return ans
}

// 对符合要求的点进行DFS
func DFS(grid [][]byte, x, y int) {
	grid[x][y] = '0' // 将陆地变为海洋，防止走回头路
	for i := 0; i < len(dx); i++ {
		nx, ny := x+dx[i], y+dy[i]
		if nx < 0 || ny < 0 || nx >= len(grid) || ny >= len(grid[nx]) || grid[nx][ny] == '0' {
			continue
		}
		DFS(grid, nx, ny)
	}
}

/*
	题目链接:
		https://leetcode-cn.com/problems/number-of-islands/							岛屿数量
		https://leetcode-cn.com/problems/number-of-closed-islands/submissions/		统计封闭岛屿的数目
		https://leetcode-cn.com/problems/number-of-enclaves/						飞地的数量
*/

/*
	总结
	1. 搜索一般模板:
			(1) 确定寻找方向(一般是上下左右)	---->    我们可以采用dx、dy数组
			(2) 防止走回头路					---->	 我们可以修改原矩阵的信息、或者创建一个isVisited映射
			(3) 判断该点是否可以走				---->    可以写一个judge函数，如果该点走过或者越界则返回false
*/
