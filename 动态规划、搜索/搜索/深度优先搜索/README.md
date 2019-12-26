

@[toc]
# 二维矩阵深度优先搜索 (TODO)
## 1. 框架
```go
var dx []int    // x变化向量
var dy []int    // y变化向量

// DFS的调用者
func DFSCaller() int {
	/* 1. 确定搜索方向dx、dy  */
	/* 2. 对坐标进行DFS搜索 (这里开始调用DFS函数) */
}

// DFS递归搜索
func DFS(grid [][]int, x int, y int) {
	/* 3. 判断进行DFS的坐标是否符合要求 */
    /* 4. 对坐标进行标记，防止走回头路 */
	for /* 5.遍历所有方向 */ {
		/* 6. 对下一个坐标进行DFS搜索 */
	}
}
```

## 2. 实例
### 2.1 岛屿数量
```go
/*
	给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，
	并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。

	题目链接: https://leetcode-cn.com/problems/number-of-islands/	岛屿数量
*/

var dx []int    // x变化向量
var dy []int    // y变化向量

// 这个就是DFSCaller
func numIslands(grid [][]byte) int {
	/* 1. 确定搜索方向 */
	dx = []int{0, 0, 1, -1}
	dy = []int{1, -1, 0, 0}
	
	/* 2. 对坐标进行DFS搜索 (这里开始调用DFS函数) */
	ans := 0
	for i := 0; i < len(grid); i++ {
		for t := 0; t < len(grid[i]); t++ {
			if grid[i][t] == '1'{
				ans++
			}
			DFS(grid, i, t)
		}
	}
	
	return ans
}

// DFS递归搜索
func DFS(grid [][]byte, x, y int) {
	/* 3. 判断进行DFS的坐标是否符合要求 */
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[x]) || grid[x][y] == '0' {
	    // 不满足时
	    return
	}
	// 满足时
  	/* 4. 对坐标进行标记，防止走回头路 */
	grid[x][y] = '0' // 将陆地变为海洋，防止走回头路
	
	/* 5.遍历所有方向 */
	for i := 0; i < len(dx); i++ {
		
		/* 6. 对下一个坐标进行DFS搜索 */
		nx, ny := x+dx[i], y+dy[i]
		DFS(grid, nx, ny)
	}
}
```
### 2.2 飞地的数量

```go
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
```
## 3. 练习题
- [ ] [200. 岛屿数量](https://leetcode-cn.com/problems/number-of-islands/)
- [ ] [1020. 飞地的数量](https://leetcode-cn.com/problems/number-of-enclaves/submissions/)
