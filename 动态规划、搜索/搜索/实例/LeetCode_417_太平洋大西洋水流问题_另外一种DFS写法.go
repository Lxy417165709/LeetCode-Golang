package main

var pacific [155][155]bool
var atlantic [155][155]bool
var dx []int
var dy []int
var isVisit map[int]bool

const (
	inf = 100000000000
)

func hash(x, y int) int {
	return (x << 10) | y
}

// 判断接下来的点是否可以走(必须保证不走回头路)
func judge(matrix [][]int, x, y int, preHight int) bool {
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
	// preHight 表示先前的山脉高度，如果先前的小于当前的，那么这个点就可以走
	return preHight <= matrix[x][y]
}

// occenFlag == 0 表示是太平洋，occenFlag == 1是大西洋
func DFS(matrix [][]int, x, y int, occenFlag int, preHight int) {
	if !judge(matrix, x, y, preHight) {
		return
	}
	isVisit[hash(x, y)] = true
	if occenFlag == 0 {
		pacific[x][y] = true
	}
	if occenFlag == 1 {
		atlantic[x][y] = true
	}
	for i := 0; i < len(dx); i++ {
		DFS(matrix, x+dx[i], y+dy[i], occenFlag, matrix[x][y])
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

/*
	题目链接:
		https://leetcode-cn.com/problems/pacific-atlantic-water-flow/			太平洋大西洋水流问题
*/
/*
	总结
	1. 这个代码和之前的DFS差别在于，这份代码是在DFS进入时判断点的合法性，而先前的那一份是在DFS前判断点的合法性。
*/
