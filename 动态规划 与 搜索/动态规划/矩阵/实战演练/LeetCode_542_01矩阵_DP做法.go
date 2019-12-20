package main

const (
	inf = 1000000000
)

func updateMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return matrix
	}
	m, n := len(matrix), len(matrix[0])
	// 先把"1"赋值为inf,表示"1"到"0"的距离此时为无穷大
	for i := 0; i < m; i++ {
		for t := 0; t < n; t++ {
			if matrix[i][t] == 1 {
				matrix[i][t] = inf
			}
		}
	}
	// 从上到下、从左到右遍历
	for i := 0; i < m; i++ {
		for t := 0; t < n; t++ {
			if i > 0 {
				matrix[i][t] = min(matrix[i][t], matrix[i-1][t]+1)
			}
			if t > 0 {
				matrix[i][t] = min(matrix[i][t], matrix[i][t-1]+1)
			}
		}
	}
	// 从下到上、从右到左遍历
	for i := m - 1; i >= 0; i-- {
		for t := n - 1; t >= 0; t-- {
			if i < m-1 {
				matrix[i][t] = min(matrix[i][t], matrix[i+1][t]+1)
			}
			if t < n-1 {
				matrix[i][t] = min(matrix[i][t], matrix[i][t+1]+1)
			}
		}
	}
	return matrix
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*
	总结
	1. 这题目之前是采用BFS做的，其实也可以用DP做。
	2. 思路如下:
			(1) 由于"1"到"0"的最短距离取决于"1"的四周，于是可以得出
					matrix[i][t] = min(matrix[i][t+1]+1, matrix[i][t+1]+1, matrix[i-1][t]+1 matrix[i][t-1]+1)
				这里涉及了4个方向，于是我们分2次遍历，第一次从上到下、从左到右遍历，第二次从下到上、从右到左遍历
				综合以上的遍历结果，就能正确的得出matrix[i][t]的值了。
			(2) 详情看代码
*/
