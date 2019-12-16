package main

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	l, r := 0, m*n-1
	for l <= r {
		mid := (l + r) / 2
		if matrix[mid/n][mid%n] == target {
			return true
		} else {
			if matrix[mid/n][mid%n] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return false
}

/*
	总结
	1. 这题就是把一个一维递增数组变为二维递增矩阵，那么我们此时也可以采用二分查找，只是要做一些相应变换。
*/
