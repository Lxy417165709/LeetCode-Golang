package matrix_util

// GetHeightAndWidth 获取矩阵高度、宽度。
func GetHeightAndWidth(matrix [][]int) (int, int) {
	if len(matrix) == 0 {
		return 0, 0
	}
	return len(matrix), len(matrix[0])
}

// GetHeightAndWidthOfByteMatrix 获取矩阵高度、宽度。
func GetHeightAndWidthOfByteMatrix(matrix [][]byte) (int, int) {
	if len(matrix) == 0 {
		return 0, 0
	}
	return len(matrix), len(matrix[0])
}
