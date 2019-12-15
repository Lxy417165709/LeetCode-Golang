package main

const (
	queenFlag = 10 // 地图上皇后的标识
)

var ans [][]int
var dx [8]int
var dy [8]int

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	// 构建国王坐标的变化向量
	dx = [8]int{0, 0, -1, -1, -1, 1, 1, 1}
	dy = [8]int{-1, 1, 0, -1, 1, 0, -1, 1}

	ans = make([][]int, 0)
	grid := [8][8]int{}
	for i := 0; i < len(queens); i++ {
		grid[queens[i][0]][queens[i][1]] = queenFlag
	}
	for t := 0; t < len(dx); t++ {
		DFS(grid, king[0], king[1], t)
	}
	return ans
}

// 判断该点是否可走
func judge(grid [8][8]int, x, y int) bool {
	if len(grid) == 0 {
		return false
	}
	m, n := len(grid), len(grid[0])
	if x < 0 || y < 0 || x >= m || y >= n {
		return false
	}
	return true
}

// DFS扫描 (由于国王每次只能向一个方向走，所以我这里加了个direction参数，表示国王此时的移动方向)
func DFS(grid [8][8]int, x int, y int, direction int) {

	if judge(grid, x, y) == false {
		return
	}
	if grid[x][y] == queenFlag {
		ans = append(ans, []int{x, y})
		return
	}
	DFS(grid, x+dx[direction], y+dy[direction], direction)
}


/*
	总结
	1. 这题目我从国外出发，往8个方向找皇后，找到皇后就停止扫描，并把皇后的坐标加入结果集。
*/
