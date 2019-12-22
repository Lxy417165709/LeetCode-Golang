package main

func findDiagonalOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	m, n := len(matrix), len(matrix[0])
	ans := make([]int, 0, n*m)
	curX, curY, indexSum := 0, 0, 0 // indexSum = curX + curY，在同一对角线上遍历时indexSum是不变的
	// n + m - 1: 这是最后一条对角线的索引和
	for indexSum != n+m-1 {
		if indexSum%2 == 0 {
			curX = min(indexSum, m-1)
			curY = indexSum - curX
		} else {
			curY = min(indexSum, n-1)
			curX = indexSum - curY
		}
		for curX >= 0 && curX < m && curY >= 0 && curY < n {
			ans = append(ans, matrix[curX][curY])
			if indexSum%2 == 0 {
				curX--
				curY = indexSum - curX
			} else {
				curY--
				curX = indexSum - curY
			}
		}
		indexSum++
	}
	return ans
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*
	总结
	1. 优美了一点。
*/
