package main

var dx []int
var dy []int

// DFSCaller
func numEnclaves(A [][]int) int {

	dx = []int{0, 0, -1, 1}
	dy = []int{-1, 1, 0, 0}

	m, n := len(A), len(A[0])

	// 去除能到达边界的陆地
	for i := 0; i < len(A); i++ {
		DFS(A, i, 0)
		DFS(A, i, n-1)
	}
	for i := 0; i < len(A[0]); i++ {
		DFS(A, 0, i)
		DFS(A, m-1, i)
	}

	ans := 0
	for i := 0; i < m; i++ {
		for t := 0; t < n; t++ {
			if i < 0 || t < 0 || i >= m || t >= n || A[i][t] == 0 {
				continue
			}
			ans++
		}
	}
	return ans
}

// DFS
func DFS(grid [][]int, x, y int) {

	// 这里加了个判断条件，是为了阻止不符合要求的节点进行DFS。
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[x]) || grid[x][y] == 0{
		return
	}

	grid[x][y] = 0
	for i := 0; i < len(dx); i++ {
		nx, ny := x+dx[i], y+dy[i]
		if nx < 0 || ny < 0 || nx >= len(grid) || ny >= len(grid[nx]) || grid[nx][ny] == 0 {
			continue
		}
		DFS(grid, x+dx[i], y+dy[i])
	}
}
