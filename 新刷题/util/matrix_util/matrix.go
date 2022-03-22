package matrix_util

// SpiralOrder 获取矩阵螺旋顺。
func SpiralOrder(matrix [][]int) []int {
	// 1. 获取矩阵高度、宽度。
	height, width := GetHeightAndWidth(matrix)

	// 2. 初始化。
	upBound := 0
	downBound := height - 1
	leftBound := 0
	rightBound := width - 1

	// 3. 只有一个元素时。
	if upBound == downBound && leftBound == rightBound {
		return matrix[0]
	}

	// 4. 螺旋遍历。
	result := make([]int, 0)
	for {
		// 4.1 左->右遍历。
		for i := leftBound; i <= rightBound; i++ {
			result = append(result, matrix[upBound][i])
		}
		upBound++
		if upBound > downBound {
			break
		}

		// 4.2 上->下遍历。
		for i := upBound; i <= downBound; i++ {
			result = append(result, matrix[i][rightBound])
		}
		rightBound--
		if leftBound > rightBound {
			break
		}

		// 4.3 右->左遍历。
		for i := rightBound; i >= leftBound; i-- {
			result = append(result, matrix[downBound][i])
		}
		downBound--
		if upBound > downBound {
			break
		}

		// 4.4 下->上遍历。
		for i := downBound; i >= upBound; i-- {
			result = append(result, matrix[i][leftBound])
		}
		leftBound++
		if leftBound > rightBound {
			break
		}
	}

	// 5. 返回。
	return result
}

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
