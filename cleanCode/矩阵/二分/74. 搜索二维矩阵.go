package 二分


// ----------------------------------- 方法1: 暴力 -----------------------------------
// 执行用时：4 ms,   在所有 Go 提交中击败了 99.07%  的用户
// 内存消耗：3.8 MB, 在所有 Go 提交中击败了 100.00% 的用户
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
// 执行用时：8 ms,   在所有 Go 提交中击败了  84.49% 的用户
// 内存消耗：3.8 MB, 在所有 Go 提交中击败了 100.00% 的用户
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
// ----------------------------------- 方法3: 以矩阵为整体，进行二分查找 -----------------------------------
// 执行用时：8 ms,   在所有 Go 提交中击败了 84.49%  的用户
// 内存消耗：3.8 MB, 在所有 Go 提交中击败了 100.00% 的用户
func searchMatrix(matrix [][]int, target int) bool {
	rows, cols := getRowsAndCols(matrix)
	leftIndex,rightIndex := 0,rows*cols-1
	for leftIndex<=rightIndex{
		midIndex := leftIndex+(rightIndex-leftIndex)/2
		midValue := getValue(matrix,midIndex)
		if midValue==target{
			return true
		}
		if midValue>target{
			rightIndex = midIndex-1
		}else{
			leftIndex = midIndex+1
		}
	}
	return false
}

func getValue(matrix [][]int,index int) int{
	_,cols := getRowsAndCols(matrix)
	return matrix[index/cols][index%cols]
}

func getRowsAndCols(matrix [][]int) (int, int) {
	if len(matrix) == 0 {
		return 0, 0
	}
	return len(matrix), len(matrix[0])
}
// ----------------------------------- 方法4: 从右上角开始查找 -----------------------------------
// 执行用时：8 ms,   在所有 Go 提交中击败了 84.49%  的用户
// 内存消耗：3.8 MB, 在所有 Go 提交中击败了 100.00% 的用户
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

