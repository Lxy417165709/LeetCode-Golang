package main

var dx []int
var dy []int

func judge(board [][]byte, x, y int) bool {
	if len(board) == 0 {
		return false
	}
	m, n := len(board), len(board[0])
	if x < 0 || y < 0 || x >= m || y >= n {
		return false
	}
	if board[x][y] == 'X' || board[x][y] == 'S' {
		return false
	}
	return true
}

func DFS(board [][]byte, x, y int) {
	if !judge(board, x, y) {
		return
	}
	board[x][y] = 'S'
	for i := 0; i < len(dx); i++ {
		DFS(board, x+dx[i], y+dy[i])
	}
}

func solve(board [][]byte) {
	dx = []int{0, 0, -1, 1}
	dy = []int{1, -1, 0, 0}
	if len(board) == 0 {
		return
	}
	m, n := len(board), len(board[0])
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