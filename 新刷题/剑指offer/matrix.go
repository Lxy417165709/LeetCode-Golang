package 剑指offer

// getHeightAndWidth 获取矩阵高度、宽度。
func getHeightAndWidth(matrix [][]int) (int, int) {
	if len(matrix) == 0 {
		return 0, 0
	}
	return len(matrix), len(matrix[0])
}
