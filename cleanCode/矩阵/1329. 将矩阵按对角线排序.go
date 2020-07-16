package main

// ------------------------------- 基于快速排序的矩阵对角线排序 (原地排序) -------------------------------
func diagonalSort(mat [][]int) [][]int {
	if len(mat) == 0 {
		return [][]int{}
	}
	rows, cols := getRowsAndCols(mat)
	for i := 0; i < rows; i++ {
		leftX, leftY := i, 0
		rightX, rightY := getRightXAndRightY(mat, leftX, leftY)
		quickSortForMatrixDiagonal(mat, leftX, leftY, rightX, rightY)
	}
	for i := 0; i < cols; i++ {
		leftX, leftY := 0, i
		rightX, rightY := getRightXAndRightY(mat, leftX, leftY)
		quickSortForMatrixDiagonal(mat, leftX, leftY, rightX, rightY)
	}
	return mat
}

func getRightXAndRightY(mat [][]int, leftX, leftY int) (int, int) {
	rows, cols := getRowsAndCols(mat)
	distanceToRightBound := cols - leftY - 1
	distanceToDownBound := rows - leftX - 1
	minDistanceToBound := min(distanceToDownBound, distanceToRightBound)
	return leftX + minDistanceToBound, leftY + minDistanceToBound

	//  这个函数这样写也可以
	//  for {
	//		if !isValidCoordinate(mat, leftX+1, leftY+1) {
	//			return leftX,leftY
	//		}
	//		leftX++
	//		leftY++
	//	}
}

func quickSortForMatrixDiagonal(mat [][]int, leftX, leftY, rightX, rightY int) {
	if !isValidCoordinate(mat, leftX, leftY) || !isValidCoordinate(mat, rightX, rightY) {
		return
	}
	if leftX > rightX || leftY > rightY {
		return
	}
	partitionX, partitionY := partitionForMatrixDiagonal(mat, leftX, leftY, rightX, rightY)
	quickSortForMatrixDiagonal(mat, leftX, leftY, partitionX-1, partitionY-1)
	quickSortForMatrixDiagonal(mat, partitionX+1, partitionY+1, rightX, rightY)
}

func partitionForMatrixDiagonal(mat [][]int, leftX, leftY, rightX, rightY int) (int, int) {
	guardX, guardY := leftX, leftY
	guardNum := mat[guardX][guardY]
	for leftX <= rightX && leftY <= rightY {
		for leftX <= rightX && leftY <= rightY && mat[leftX][leftY] <= guardNum {
			leftX++
			leftY++
		}
		for leftX <= rightX && leftY <= rightY && mat[rightX][rightY] >= guardNum {
			rightX--
			rightY--
		}
		if leftX <= rightX && leftY <= rightY {
			mat[rightX][rightY], mat[leftX][leftY] = mat[leftX][leftY], mat[rightX][rightY]
		}
	}
	mat[rightX][rightY], mat[guardX][guardY] = mat[guardX][guardY], mat[rightX][rightY]
	return rightX, rightY
}

func isValidCoordinate(mat [][]int, x, y int) bool {
	rows, cols := getRowsAndCols(mat)
	return x >= 0 && x <= rows-1 && y >= 0 && y <= cols-1
}

func getRowsAndCols(mat [][]int) (int, int) {
	return len(mat), len(mat[0])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*
	题目链接: https://leetcode-cn.com/problems/sort-the-matrix-diagonally/
	总结:
		1. 这题我基于快速排序，对矩阵对角线进行排序。
*/
