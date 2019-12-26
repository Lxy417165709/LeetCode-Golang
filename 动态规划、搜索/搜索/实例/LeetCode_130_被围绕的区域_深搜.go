package main

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

/*
	题目链接:
		https://leetcode-cn.com/problems/surrounded-regions/solution/		被围绕的区域
*/
/*
	总结
	1. 这题的思路是:
			(1) 扫描矩阵边界，利用DFS将边界岛屿的每一块变为'S'。	(字符不唯一，可以选择其他字符，只要不是'O','X')
			(2) 扫描整个矩阵，将'O'变'X', 'S'变'O'
			(3) 完毕~
	2. 个人感觉DFS一般比较容易编写~
*/
