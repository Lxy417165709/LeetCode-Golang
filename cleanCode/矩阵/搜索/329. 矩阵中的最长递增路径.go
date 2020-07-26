package 搜索

import "sort"

// ------------------------------ 记忆化搜索 ------------------------------
const INF = 1000000000

var vector = [][]int{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}
var longestPathStartWith [][]int

func longestIncreasingPath(matrix [][]int) int {
	rows, cols := getRowsAndCols(matrix)
	longestPathStartWith = get2DSlice(rows, cols)
	maxLength := 0
	for i := 0; i < rows; i++ {
		for t := 0; t < cols; t++ {
			maxLength = max(maxLength, getLongestIncreasingPath(matrix, i, t))
		}
	}
	return maxLength
}

func get2DSlice(rows, column int) [][]int {
	slice := make([][]int, rows)
	for i := 0; i < len(slice); i++ {
		slice[i] = make([]int, column)
	}
	return slice
}

func getRowsAndCols(matrix [][]int) (int, int) {
	if len(matrix) == 0 {
		return 0, 0
	}
	return len(matrix), len(matrix[0])
}

func getLongestIncreasingPath(matrix [][]int, startX, startY int) int {
	rows, cols := getRowsAndCols(matrix)
	if longestPathStartWith[startX][startY] != 0 {
		return longestPathStartWith[startX][startY]
	}
	maxLength := 1
	for i := 0; i < len(vector); i++ {
		nextX, nextY := startX+vector[i][0], startY+vector[i][1]
		if nextX < 0 || nextY < 0 || nextX >= rows || nextY >= cols || matrix[startX][startY] <= matrix[nextX][nextY] {
			continue
		}
		maxLength = max(
			maxLength,
			getLongestIncreasingPath(matrix, nextX, nextY)+1,
		)
	}
	longestPathStartWith[startX][startY] = maxLength
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ----------------------------- 动态规划 --------------------------------
const INF = 1000000000

var vector [][]int = [][]int{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}
	units := getUnits(matrix)
	longestPathOverWith := getLongestPathOverWith(matrix, units)
	return getMax(longestPathOverWith)
}

func getUnits(matrix [][]int) []Unit {
	rows, cols := getRowsAndCols(matrix)
	units := make([]Unit, 0)
	for i := 0; i < rows; i++ {
		for t := 0; t < cols; t++ {
			units = append(units, Unit{matrix[i][t], i, t})
		}
	}
	return units
}

func getLongestPathOverWith(matrix [][]int, units []Unit) [][]int {
	rows, cols := getRowsAndCols(matrix)
	longestPathOverWith := get2DSlice(rows, cols)
	sort.Slice(units, func(i, t int) bool {
		return units[i].Value < units[t].Value
	})

	for i := 0; i < len(units); i++ {
		overX, overY := units[i].X, units[i].Y
		longestPathOverWith[overX][overY] = 1
		for t := 0; t < len(vector); t++ {
			preOverX, preOverY := overX+vector[t][0], overY+vector[t][1]
			if preOverX < 0 || preOverY < 0 || preOverX >= rows || preOverY >= cols {
				continue
			}
			if matrix[overX][overY] <= matrix[preOverX][preOverY] {
				continue
			}
			longestPathOverWith[overX][overY] = max(
				longestPathOverWith[overX][overY],
				longestPathOverWith[preOverX][preOverY]+1,
			)
		}
	}
	return longestPathOverWith
}

func getMax(matrix [][]int) int {
	mx := -INF
	for i := 0; i < len(matrix); i++ {
		for t := 0; t < len(matrix[i]); t++ {
			mx = max(mx, matrix[i][t])
		}
	}
	return mx
}

func get2DSlice(rows, column int) [][]int {
	slice := make([][]int, rows)
	for i := 0; i < len(slice); i++ {
		slice[i] = make([]int, column)
	}
	return slice
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getRowsAndCols(matrix [][]int) (int, int) {
	if len(matrix) == 0 {
		return 0, 0
	}
	return len(matrix), len(matrix[0])
}

// ----------------------- Unit -----------------------
type Unit struct {
	Value int
	X     int
	Y     int
}

/*
	题目链接: https://leetcode-cn.com/problems/longest-increasing-path-in-a-matrix/
	总结:
		1. 该题能防止走回头路,关键在于： matrix[startX][startY] <= matrix[nextX][nextY]
		2. 官方还有使用拓扑排序AC的。 (上面的动态规划就是运用了拓扑排序的思维，要保证值小的先被求出)
*/
