package main

var dx []int    // x变化向量
var dy []int    // y变量向量

// DFS的调用者
func numEnclaves(A [][]int) int {

	/* 1. 确定搜索方向dx、dy */
	dx = []int{0, 0, -1, 1}
	dy = []int{-1, 1, 0, 0}

	m, n := len(A), len(A[0])

	/* 2. 对坐标进行DFS搜索 (这里开始调用DFS函数) */
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

	/* 3. 判断进行DFS的坐标是否符合要求 */
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[x]) || grid[x][y] == 0{
		return
	}

	/* 4. 对坐标进行标记，防止走回头路 */
	grid[x][y] = 0

	/* 5.遍历所有方向 */
	for i := 0; i < len(dx); i++ {

		/* 6. 对下一个坐标进行DFS搜索 */
		nx, ny := x+dx[i], y+dy[i]
		DFS(grid, nx, ny)
	}
}