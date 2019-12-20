package main

var pacific [155][155]bool
var atlantic [155][155]bool
var dx []int
var dy []int
var isVisit map[int]bool

func hash(x, y int) int {
	return (x << 10) | y
}

// 判断接下来的点是否可以走(必须保证不走回头路)
func judge(matrix [][]int, x, y int, nowHight int) bool {
	if len(matrix) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	if x < 0 || y < 0 || x >= m || y >= n {
		return false
	}
	if isVisit[hash(x, y)] {
		return false
	}
	return nowHight <= matrix[x][y] // 由于是使用逆流，所以此时是需要保证低的流到高的。
}

// occenFlag == 0 表示是太平洋，occenFlag == 1是大西洋
func DFS(matrix [][]int, x, y int, occenFlag int) {
	isVisit[hash(x, y)] = true
	if occenFlag == 0 {
		pacific[x][y] = true
	}
	if occenFlag == 1 {
		atlantic[x][y] = true
	}
	for i := 0; i < len(dx); i++ {
		nx, ny := x+dx[i], y+dy[i]
		if judge(matrix, nx, ny, matrix[x][y]) {
			DFS(matrix, nx, ny, occenFlag)
		}
	}
}

func pacificAtlantic(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return [][]int{}
	}
	m, n := len(matrix), len(matrix[0])
	pacific = [155][155]bool{}  // pacific[x][y]  表示坐标(x,y)能否到达太平洋
	atlantic = [155][155]bool{} // atlantic[x][y] 表示坐标(x,y)能否到达大西洋
	dx = []int{1, -1, 0, 0}
	dy = []int{0, 0, 1, -1}

	// 找能到达太平洋的坐标
	isVisit = make(map[int]bool)
	for i := 0; i < m; i++ {
		DFS(matrix, i, 0, 0)
	}
	for t := 0; t < n; t++ {
		DFS(matrix, 0, t, 0)
	}

	// 找能到达大西洋的坐标
	isVisit = make(map[int]bool)
	for i := 0; i < m; i++ {
		DFS(matrix, i, n-1, 1)
	}
	for t := 0; t < n; t++ {
		DFS(matrix, m-1, t, 1)
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

/*
	题目链接:
		https://leetcode-cn.com/problems/pacific-atlantic-water-flow/			太平洋大西洋水流问题
*/
/*
	总结
	1. 这题目采用逆向思维，从边界出发，找能到达两个大洋的山脉。而不是从中心出发，判断该点能否到达两个大洋。
*/
