package main

// 实现矩阵的转置
func transpose(A [][]int) [][]int {
	if len(A) == 0 {
		return A
	}
	m, n := len(A), len(A[0])
	B := [][]int{}
	for i := 0; i < n; i++ {
		B = append(B, make([]int, m))
	}
	for i := 0; i < m; i++ {
		for t := 0; t < n; t++ {
			B[t][i] = A[i][t]
		}
	}
	return B
}

/*
	题目链接:
		https://leetcode-cn.com/problems/transpose-matrix/solution/zhuan-zhi-ju-zhen-by-leetcode/		转置矩阵
*/
