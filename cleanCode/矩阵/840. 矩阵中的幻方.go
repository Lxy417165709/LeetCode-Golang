package main

// ----------------------------- 方法1: 暴力法 -----------------------------
func numMagicSquaresInside(grid [][]int) int {
	rows, cols := getRowsAndCols(grid)
	countOfMagicSquare := 0
	for i := 2; i < rows; i++ {
		for t := 2; t < cols; t++ {
			if isMagicSquare(grid, i, t) {
				countOfMagicSquare++
			}
		}
	}
	return countOfMagicSquare
}

func isMagicSquare(matrix [][]int, x, y int) bool {
	countOfNum := getCountOfNumOfSquare(matrix, x, y)
	rowSum, colSum, mainDiagonalSum, paraDiagonalSum := getAllSumsOfSquare(matrix, x, y)
	return isAllDigitDifferentAndInRangeOneToNine(countOfNum) && isAllSumEqual(rowSum, colSum, mainDiagonalSum, paraDiagonalSum)
}

func getCountOfNumOfSquare(matrix [][]int, x, y int) map[int]int {
	countOfNum := make(map[int]int)
	for i := x - 2; i <= x; i++ {
		for t := y - 2; t <= y; t++ {
			countOfNum[matrix[i][t]]++
		}
	}
	return countOfNum
}

func getAllSumsOfSquare(matrix [][]int, x, y int) ([]int, []int, int, int) {
	rowSum, colSum := make([]int, 3), make([]int, 3)
	mainDiagonalSum, paraDiagonalSum := 0, 0
	for i := x - 2; i <= x; i++ {
		for t := y - 2; t <= y; t++ {
			relativeX, relativeY := i-x+2, t-y+2
			rowSum[relativeX] += matrix[i][t]
			colSum[relativeY] += matrix[i][t]
			if relativeX == relativeY {
				mainDiagonalSum += matrix[i][t]
			}
			if relativeX+relativeY == 2 {
				paraDiagonalSum += matrix[i][t]
			}
		}
	}
	return rowSum, colSum, mainDiagonalSum, paraDiagonalSum
}

func isAllDigitDifferentAndInRangeOneToNine(countOfNum map[int]int) bool {
	for i := 1; i <= 9; i++ {
		if countOfNum[i] != 1 {
			return false
		}
	}
	return true
}

func isAllSumEqual(rowSum, colSum []int, mainDiagonalSum, paraDiagonalSum int) bool {
	refSum := mainDiagonalSum
	if paraDiagonalSum != refSum {
		return false
	}
	for i := 0; i < len(rowSum); i++ {
		if refSum != rowSum[i] {
			return false
		}
	}
	for i := 0; i < len(colSum); i++ {
		if refSum != colSum[i] {
			return false
		}
	}
	return true
}

func getRowsAndCols(matrix [][]int) (int, int) {
	return len(matrix), len(matrix[0])
}

// ----------------------------- 方法2: 暴力法优化方向(无代码) -----------------------------
// 1. 暴力法的求取总和，可以使用前缀和进行优化。
// 2. 上面的 countOfNum 类型是 map[int]int, 其实可以替换为长度为 9 的数组。


// ----------------------------- 方法3: 搜索 + 字典树(无代码)  -----------------------------
// 1. 可以通过搜索的方式，得出所有 3 阶幻方。 (看评论区说只有 8 个)
// 2. 将 3 阶幻方矩阵转换为数组，那这个数组就相当于字符串。于是可以以这 8 个幻方为基础，建立字典树。
// 3. 回到题目，暴力遍历 grid 时，通过字典树，就能判断 grid 中每一个 3X3 的区域是否为幻方了。

/*
	题目链接:
	总结;
		1. 这题其实可以求出所有幻方，将其构建为一颗字典树，这样判断是否是幻方就很简单了。
		2. 这题我采用了暴力。
		3. 可以用前缀和优化上面的代码。
*/
