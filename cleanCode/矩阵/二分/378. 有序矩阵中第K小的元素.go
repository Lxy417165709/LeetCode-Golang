package 二分

// ------------------------- 方法1: 随机选择算法 -------------------------
// 这个方法适用于任何矩阵，不要求矩阵有序。
const flagOfCannotFind = -1

func kthSmallest(matrix [][]int, k int) int {
	rows, cols := getRowsAndCols(matrix)
	return getKthSmallest(matrix, 0, rows*cols-1, k)
}

func getKthSmallest(matrix [][]int, left, right, k int) int {
	if left > right {
		return flagOfCannotFind
	}

	indexOfRightPosition := partition(matrix, left, right)
	xthSmall := indexOfRightPosition + 1
	if xthSmall == k {
		return getValue(matrix, indexOfRightPosition)
	}
	if xthSmall > k {
		return getKthSmallest(matrix, left, indexOfRightPosition-1, k)
	} else {
		return getKthSmallest(matrix, indexOfRightPosition+1, right, k)
	}
}

func partition(matrix [][]int, left, right int) int {
	guardIndex := left
	guardNum := getValue(matrix, guardIndex)
	for left <= right {
		if left <= right && getValue(matrix, left) <= guardNum {
			left++
		}
		if left <= right && getValue(matrix, right) >= guardNum {
			right--
		}
		if left <= right {
			swap(matrix, left, right)
		}
	}
	swap(matrix, right, guardIndex)
	return right
}

func getValue(matrix [][]int, index int) int {
	rows, _ := getRowsAndCols(matrix)
	return matrix[index/rows][index%rows]
}

func swap(matrix [][]int, index1, index2 int) {
	rows, _ := getRowsAndCols(matrix)
	matrix[index1/rows][index1%rows], matrix[index2/rows][index2%rows] = matrix[index2/rows][index2%rows], matrix[index1/rows][index1%rows]
}

func getRowsAndCols(matrix [][]int) (int, int) {
	return len(matrix), len(matrix[0])
}

// ------------------------- 方法2: 插入排序(超时) -------------------------
// 这个方法适用于任何矩阵，不要求矩阵有序。但是有序的话，插入排序会比较快。 (很尴尬的是超时了)
func kthSmallest(matrix [][]int, k int) int {
	insertSort(matrix)
	return getValue(matrix, k-1)
}

func insertSort(matrix [][]int) {
	rows, cols := getRowsAndCols(matrix)
	left, right := 1, rows*cols-1
	for i := left; i <= right; i++ {
		for t := i; t >= left; t-- {
			if getValue(matrix, t) < getValue(matrix, t-1) {
				swap(matrix, t, t-1)
			} else {
				break
			}
		}
	}
}

func getValue(matrix [][]int, index int) int {
	rows, _ := getRowsAndCols(matrix)
	return matrix[index/rows][index%rows]
}

func swap(matrix [][]int, index1, index2 int) {
	rows, _ := getRowsAndCols(matrix)
	matrix[index1/rows][index1%rows], matrix[index2/rows][index2%rows] = matrix[index2/rows][index2%rows], matrix[index1/rows][index1%rows]
}

func getRowsAndCols(matrix [][]int) (int, int) {
	return len(matrix), len(matrix[0])
}

// ------------------------- 方法3: 二分 -------------------------
// 这个方法需要每行、每列非递减。
func kthSmallest(matrix [][]int, k int) int {
	leftNum, rightNum := getMinNumAndMaxNum(matrix)
	for leftNum <= rightNum {
		midNum := leftNum + (rightNum-leftNum)/2
		countOfNumLessOrEqualMidNum := getCountOfNumLessOrEqual(matrix, midNum)
		if countOfNumLessOrEqualMidNum == k {
			rightNum = midNum - 1
			// 往左边寻找，因为 midNum 可能不在矩阵中。
			// 以题目示例为例，14、13都是第8小，但 14 不在矩阵中，13 在。
			// 个人感觉，这句代码和 getCountOfNumLessOrEqual 有关。
		} else {
			if countOfNumLessOrEqualMidNum > k {
				rightNum = midNum - 1
			} else {
				leftNum = midNum + 1
			}
		}
	}
	return leftNum
}
func getMinNumAndMaxNum(matrix [][]int) (int, int) {
	rows, cols := getRowsAndCols(matrix)
	return matrix[0][0], matrix[rows-1][cols-1]
}

func getCountOfNumLessOrEqual(matrix [][]int, ref int) int {
	rows, cols := getRowsAndCols(matrix)
	curX, curY := rows-1, 0
	countOfNumLessOrEqual := 0
	for curX >= 0 && curY <= cols-1 {
		if matrix[curX][curY] <= ref {
			countOfNumLessOrEqual += curX + 1
			curY++
		} else {
			curX--
		}
	}
	return countOfNumLessOrEqual
}

func getRowsAndCols(matrix [][]int) (int, int) {
	return len(matrix), len(matrix[0])
}

// ------------------------- 方法4: 快速排序、归并排序、堆(无代码) -------------------------
// 略

/*
	题目链接: https://leetcode-cn.com/problems/kth-smallest-element-in-a-sorted-matrix/
	总结:
		1. 官方题目有用二分的
		2. 也有用归并排序的
*/
