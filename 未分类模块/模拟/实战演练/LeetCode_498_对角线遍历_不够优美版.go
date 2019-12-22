package main

func findDiagonalOrder(matrix [][]int) []int {

	if len(matrix) == 0 {
		return []int{}
	}
	m, n := len(matrix), len(matrix[0])
	times := (n + m) / 2 // 循环次数
	ans := make([]int, 0, n*m)
	beginx, beginy, base := 0, 0, 0
	for times != 0 {
		times--

		// 向左上遍历
		beginx = min(base, m-1)
		beginy = base - beginx
		for beginx >= 0 && beginx < m && beginy >= 0 && beginy < n {
			ans = append(ans, matrix[beginx][beginy])
			beginx--
			beginy = base - beginx
		}
		base++
		// 向右下遍历
		beginy = min(base, n-1)
		beginx = base - beginy
		for beginx >= 0 && beginx < m && beginy >= 0 && beginy < n {
			ans = append(ans, matrix[beginx][beginy])
			beginy--
			beginx = base - beginy
		}
		base++
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
	1. 这个题目就是单纯的模拟，我的模拟思路是: 以一上一下为一个轮回，在遍历过程中把元素加入结果集。
	2. 这个代码还不算太优美，接下来我改一改。
*/