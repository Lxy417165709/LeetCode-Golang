package 动态规划

import "fmt"

// -------------------- 贪心动态规划(使用前缀、后缀最小值) ----------------------
const INF = 1000000000

func minFallingPathSum(arr [][]int) int {
	rows, cols := getRowsAndCols(arr)
	minSum := get2DSlice(rows, cols)
	for t := 0; t < cols; t++ {
		minSum[0][t] = arr[0][t]
	}
	for i := 1; i < rows; i++ {
		prefixMin := getPrefixMin(minSum[i-1])
		suffixMin := getSuffixMin(minSum[i-1])
		minSum[i][0] = suffixMin[1] + arr[i][0]
		minSum[i][cols-1] = prefixMin[cols-2] + arr[i][cols-1]
		for t := 1; t < cols-1; t++ {
			minSum[i][t] = min(suffixMin[t+1], prefixMin[t-1]) + arr[i][t]
		}
	}
	return min(minSum[rows-1]...)
}

func getPrefixMin(arr []int) []int {
	prefixMin := make([]int, len(arr))
	prefixMin[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		prefixMin[i] = min(prefixMin[i-1], arr[i])
	}
	return prefixMin
}

func getSuffixMin(arr []int) []int {
	suffixMin := make([]int, len(arr))
	suffixMin[len(arr)-1] = arr[len(arr)-1]
	for i := len(arr) - 2; i >= 0; i-- {
		suffixMin[i] = min(suffixMin[i+1], arr[i])
	}
	return suffixMin
}
func getRowsAndCols(matrix [][]int) (int, int) {
	if len(matrix) == 0 {
		return 0, 0
	}
	return len(matrix), len(matrix[0])
}

func get2DSlice(rows, column int) [][]int {
	slice := make([][]int, rows)
	for i := 0; i < len(slice); i++ {
		slice[i] = make([]int, column)
	}
	return slice
}

func min(arr ...int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	a := arr[0]
	b := min(arr[1:]...)
	if a > b {
		return b
	}
	return a
}

// -------------------- 暴力动态规划 ----------------------
const INF = 1000000000

func minFallingPathSum(arr [][]int) int {
	rows, cols := getRowsAndCols(arr)
	minSum := get2DSlice(rows, cols)
	for t := 0; t < cols; t++ {
		minSum[0][t] = arr[0][t]
	}
	for i := 1; i < rows; i++ {
		for t := 0; t < cols; t++ {
			minSum[i][t] = INF
			for k := 0; k < cols; k++ {
				if t == k {
					continue
				}
				minSum[i][t] = min(minSum[i][t], minSum[i-1][k]+arr[i][t])
			}
		}
	}
	return min(minSum[rows-1]...)
}

func getRowsAndCols(matrix [][]int) (int, int) {
	if len(matrix) == 0 {
		return 0, 0
	}
	return len(matrix), len(matrix[0])
}

func get2DSlice(rows, column int) [][]int {
	slice := make([][]int, rows)
	for i := 0; i < len(slice); i++ {
		slice[i] = make([]int, column)
	}
	return slice
}

func min(arr ...int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	a := arr[0]
	b := min(arr[1:]...)
	if a > b {
		return b
	}
	return a
}

/*
	题目链接: https://leetcode-cn.com/problems/minimum-falling-path-sum-ii/submissions/
*/
