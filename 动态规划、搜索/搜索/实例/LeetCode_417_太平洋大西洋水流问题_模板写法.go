package main

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



/*
	题目链接:
		https://leetcode-cn.com/problems/pacific-atlantic-water-flow/			太平洋大西洋水流问题
*/
/*
	总结
	1. 这个代码和之前的DFS差别在于，这份代码是在DFS进入时判断点的合法性，而先前的那一份是在DFS前判断点的合法性。
*/
