package main

/*
	如果一个矩阵的每一方向由左上到右下的对角线上具有相同元素，那么这个矩阵是托普利茨矩阵。
*/

func isToeplitzMatrix(matrix [][]int) bool {
	if len(matrix) == 0 {
		return true
	}
	m, n := len(matrix), len(matrix[0])
	firstMatrix := matrix[0][:n-1]
	for i := 1; i < m; i++ {
		if !isSame(firstMatrix, matrix[i][1:]) {
			return false
		}
		firstMatrix = matrix[i][:n-1]
	}
	return true
}
// 比较两个切片是否相等(长度 + 值)
func isSame(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
/*
	题目链接:
		https://leetcode-cn.com/problems/toeplitz-matrix/submissions/		托普利茨矩阵
*/

/*
	总结
	1. 这题我采用每次比较每一行前n-1个和下一行的后n-1个，如果相同则继续判断，否则就不是托普利茨矩阵，返回false。
		比较完毕就返回true。
*/