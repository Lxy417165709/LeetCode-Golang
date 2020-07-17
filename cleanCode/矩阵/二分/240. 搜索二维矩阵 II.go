package 二分

// ----------------------------------- 方法1: 暴力 -----------------------------------
// 执行用时：32 ms,  在所有 Go 提交中击败了 77.29% 的用户
// 内存消耗：6.3 MB, 在所有 Go 提交中击败了 25.00% 的用户
func searchMatrix(matrix [][]int, target int) bool {
	rows, cols := getRowsAndCols(matrix)
	for i := 0; i < rows; i++ {
		for t := 0; t < cols; t++ {
			if matrix[i][t] == target {
				return true
			}
		}
	}
	return false
}

func getRowsAndCols(matrix [][]int) (int, int) {
	if len(matrix) == 0 {
		return 0, 0
	}
	return len(matrix), len(matrix[0])
}

// ----------------------------------- 方法2: 对每行进行二分查找 -----------------------------------
// 执行用时：32 ms,  在所有 Go 提交中击败了 77.29% 的用户
// 内存消耗：6.3 MB, 在所有 Go 提交中击败了 50.00% 的用户
func searchMatrix(matrix [][]int, target int) bool {
	rows, _ := getRowsAndCols(matrix)
	for i := 0; i < rows; i++ {
		if hasExistedInSortedArray(matrix[i], target) {
			return true
		}
	}
	return false
}

func hasExistedInSortedArray(sortedArray []int, target int) bool {
	left, right := 0, len(sortedArray)-1
	for left <= right {
		mid := (left + right) / 2
		if sortedArray[mid] == target {
			return true
		}
		if sortedArray[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

func getRowsAndCols(matrix [][]int) (int, int) {
	if len(matrix) == 0 {
		return 0, 0
	}
	return len(matrix), len(matrix[0])
}

// ----------------------------------- 方法3: 从右上角开始查找 -----------------------------------
// 执行用时：32 ms,  在所有 Go 提交中击败了 77.29% 的用户
// 内存消耗：6.3 MB, 在所有 Go 提交中击败了 50.00% 的用户
func searchMatrix(matrix [][]int, target int) bool {
	rows, cols := getRowsAndCols(matrix)
	curX, curY := 0, cols-1
	for curX <= rows-1 && curY >= 0 {
		if matrix[curX][curY] == target {
			return true
		}
		if matrix[curX][curY] > target {
			curY--
		} else {
			curX++
		}
	}
	return false
}

func getRowsAndCols(matrix [][]int) (int, int) {
	if len(matrix) == 0 {
		return 0, 0
	}
	return len(matrix), len(matrix[0])
}
