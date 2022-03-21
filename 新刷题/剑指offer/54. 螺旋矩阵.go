package 剑指offer

// spiralOrder 螺旋序遍历矩阵，获取结果。
func spiralOrder(matrix [][]int) []int {
	// 1. 获取矩阵高度、宽度。
	height, width := getHeightAndWidth(matrix)

	// 2. 初始化。
	upBound := 0
	downBound := height - 1
	leftBound := 0
	rightBound := width - 1

	// 3. 螺旋遍历。
	result := make([]int, 0)
	for {
		// 3.1 左->右遍历。
		for i := leftBound; i <= rightBound; i++ {
			result = append(result, matrix[upBound][i])
		}
		upBound++
		if upBound > downBound {
			break
		}

		// 3.2 上->下遍历。
		for i := upBound; i <= downBound; i++ {
			result = append(result, matrix[i][rightBound])
		}
		rightBound--
		if leftBound > rightBound {
			break
		}

		// 3.3 右->左遍历。
		for i := rightBound; i >= leftBound; i-- {
			result = append(result, matrix[downBound][i])
		}
		downBound--
		if upBound > downBound {
			break
		}

		// 3.4 下->上遍历。
		for i := downBound; i >= upBound; i-- {
			result = append(result, matrix[i][leftBound])
		}
		leftBound++
		if leftBound > rightBound {
			break
		}
	}

	// 4. 返回。
	return result
}
