

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
func DFS() {
	/* 3. 判断进行DFS的坐标是否符合要求 */
	/* 4. 对坐标进行标记，防止走回头路 */
	for /* 5.遍历所有方向 */ {
		/* 6. 对下一个坐标进行DFS搜索 */
	}
}
```

## 2. 实例
### 2.1 岛屿数量
[传送门](https://leetcode-cn.com/problems/number-of-islands/)
```go
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
[传送门](https://leetcode-cn.com/problems/number-of-enclaves/submissions/)
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
### 2.3 被围绕的区域
[传送门](https://leetcode-cn.com/problems/surrounded-regions/solution/)
```go
var dx []int
var dy []int

// DFS的调用者
func solve(board [][]byte) {
	/* 1. 确定搜索方向dx、dy  */
	dx = []int{0, 0, -1, 1}
	dy = []int{1, -1, 0, 0}

	if len(board) == 0 {
		return
	}
	m, n := len(board), len(board[0])

	/* 2. 对坐标进行DFS搜索 (这里开始调用DFS函数) */
	for i := 0; i < m; i++ {
		DFS(board, i, 0)
		DFS(board, i, n-1)
	}
	for t := 0; t < n; t++ {
		DFS(board, 0, t)
		DFS(board, m-1, t)
	}

	for i := 0; i < m; i++ {
		for t := 0; t < n; t++ {
			if board[i][t] == 'S' {
				board[i][t] = 'O'
				continue
			}
			if board[i][t] == 'O' {
				board[i][t] = 'X'
				continue
			}
		}
	}
}

// DFS递归搜索
func DFS(board [][]byte, x, y int) {
	/* 3. 判断进行DFS的坐标是否符合要求 */
	if x < 0 || y < 0 || x >= len(board) || y >= len(board[x]) || board[x][y] == 'X' || board[x][y] == 'S' {
		return
	}

	/* 4. 对坐标进行标记，防止走回头路 */
	board[x][y] = 'S'

	/* 5.遍历所有方向 */
	for i := 0; i < len(dx); i++ {
		/* 6. 对下一个坐标进行DFS搜索 */
		DFS(board, x+dx[i], y+dy[i])
	}
}
```


### 2.4 太平洋大西洋水流问题
[传送门](https://leetcode-cn.com/problems/pacific-atlantic-water-flow/)
```go
const inf = 100000000000
var pacific [155][155]bool
var atlantic [155][155]bool
var isVisit map[int]bool	// 用于防止走回头路

var dx []int	// x变化向量
var dy []int	// y变化向量

// DFS的调用者
func pacificAtlantic(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return [][]int{}
	}

	pacific = [155][155]bool{}  // pacific[x][y]  表示坐标(x,y)能否到达太平洋
	atlantic = [155][155]bool{} // atlantic[x][y] 表示坐标(x,y)能否到达大西洋

	/* 1. 确定搜索方向dx、dy  */
	dx = []int{1, -1, 0, 0}
	dy = []int{0, 0, 1, -1}

	m, n := len(matrix), len(matrix[0])

	/* 2. 对坐标进行DFS搜索 (这里开始调用DFS函数) */
	// 找能到达太平洋的坐标
	isVisit = make(map[int]bool)
	for i := 0; i < m; i++ {
		DFS(matrix, i, 0, 0, -inf)
	}
	for t := 0; t < n; t++ {
		DFS(matrix, 0, t, 0, -inf)
	}
	// 找能到达大西洋的坐标
	isVisit = make(map[int]bool)
	for i := 0; i < m; i++ {
		DFS(matrix, i, n-1, 1, -inf)
	}
	for t := 0; t < n; t++ {
		DFS(matrix, m-1, t, 1, -inf)
	}

	ans := [][]int{}
	for i := 0; i < m; i++ {
		for t := 0; t < n; t++ {
			if pacific[i][t] && atlantic[i][t] {
				ans = append(ans, []int{i, t})
			}
		}
	}
	return ans
}

func hash(x, y int) int {
	return (x << 10) | y
}

// DFS递归搜索
// occenFlag == 0 表示是太平洋，occenFlag == 1是大西洋
func DFS(matrix [][]int, x, y int, occenFlag int, preHight int) {
	/* 3. 判断进行DFS的坐标是否符合要求 */
	if x < 0 || y < 0 || x >= len(matrix) || y >= len(matrix[0]) || isVisit[hash(x, y)] {
		return
	}
	if preHight > matrix[x][y] {
		return
	}

	/* 4. 对坐标进行标记，防止走回头路 */
	isVisit[hash(x, y)] = true

	if occenFlag == 0 {
		pacific[x][y] = true
	}
	if occenFlag == 1 {
		atlantic[x][y] = true
	}

	/* 5.遍历所有方向 */
	for i := 0; i < len(dx); i++ {
		
		/* 6. 对下一个坐标进行DFS搜索 */
		DFS(matrix, x+dx[i], y+dy[i], occenFlag, matrix[x][y])
	}
}
```
## 3. 注意
- [ ] 以上只是DFS的大体框架，实际做题可能需要进行一些变化。

## 4. 练习题
- [ ] [130. 被围绕的区域](https://leetcode-cn.com/problems/surrounded-regions/solution/)
- [ ] [200. 岛屿数量](https://leetcode-cn.com/problems/number-of-islands/)
- [ ] [417. 太平洋大西洋水流问题](https://leetcode-cn.com/problems/pacific-atlantic-water-flow/)
- [ ] [959. 由斜杆划分区域](https://leetcode-cn.com/problems/regions-cut-by-slashes/)
- [ ] [1020. 飞地的数量](https://leetcode-cn.com/problems/number-of-enclaves/submissions/)
