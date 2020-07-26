package 单调栈

// -------------------------- 无单调栈，朴素解法 --------------------------
// 时间复杂度: O(n^2 * m^2)
func numSubmat(mat [][]int) int {
	maxRectangleHeight := getMaxRectangleHeight(mat)
	return getCountOfRectangle(maxRectangleHeight)
}

func getMaxRectangleHeight(mat [][]int) [][]int {
	rows, cols := getRowsAndCols(mat)
	maxRectangleHeight := get2DSlice(rows, cols)
	maxRectangleHeight[0][0] = mat[0][0]
	for i := 1; i < rows; i++ {
		if mat[i][0] == 1 {
			maxRectangleHeight[i][0] = maxRectangleHeight[i-1][0] + 1
		}
	}
	for t := 1; t < cols; t++ {
		maxRectangleHeight[0][t] = mat[0][t]
	}
	for i := 1; i < rows; i++ {
		for t := 1; t < cols; t++ {
			if mat[i][t] == 1 {
				maxRectangleHeight[i][t] = maxRectangleHeight[i-1][t] + 1
			}
		}
	}
	return maxRectangleHeight
}

func getCountOfRectangle(maxRectangleHeight [][]int) int {
	countOfRectangle := 0
	rows, cols := getRowsAndCols(maxRectangleHeight)
	for i := 0; i < rows; i++ {
		for t := 0; t < cols; t++ {
			for h := 1; h <= maxRectangleHeight[i][t]; h++ {
				// 纵向扩展
				countOfRectangle++
				// 横向扩展
				for leftBound := t - 1; leftBound >= 0 && maxRectangleHeight[i][leftBound] >= h; leftBound-- {
					countOfRectangle++
				}
			}
		}
	}
	return countOfRectangle
}

func getRowsAndCols(matrix [][]int) (int, int) {
	if len(matrix) == 0 {
		return 0, 0
	}
	return len(matrix), len(matrix[0])
}

func get2DSlice(rows, column int) [][]int {
	slice := make([][]int, rows)
	for i := 0; i < len(slice); i++ {
		slice[i] = make([]int, column)
	}
	return slice
}

// -------------------------- 无单调栈，朴素优化解法 --------------------------
// 时间复杂度: O(n^2 * m)
func numSubmat(mat [][]int) int {
	maxRectangleHeight := getMaxRectangleHeight(mat)
	return getCountOfRectangle(maxRectangleHeight)
}

func getMaxRectangleHeight(mat [][]int) [][]int {
	rows, cols := getRowsAndCols(mat)
	maxRectangleHeight := get2DSlice(rows, cols)
	maxRectangleHeight[0][0] = mat[0][0]
	for i := 1; i < rows; i++ {
		if mat[i][0] == 1 {
			maxRectangleHeight[i][0] = maxRectangleHeight[i-1][0] + 1
		}
	}
	for t := 1; t < cols; t++ {
		maxRectangleHeight[0][t] = mat[0][t]
	}
	for i := 1; i < rows; i++ {
		for t := 1; t < cols; t++ {
			if mat[i][t] == 1 {
				maxRectangleHeight[i][t] = maxRectangleHeight[i-1][t] + 1
			}
		}
	}
	return maxRectangleHeight
}

func getCountOfRectangle(maxRectangleHeight [][]int) int {
	countOfRectangle := 0
	rows, cols := getRowsAndCols(maxRectangleHeight)
	for i := 0; i < rows; i++ {
		for t := 0; t < cols; t++ {
			// 这里初现单调栈的端倪
			curMinHeight := maxRectangleHeight[i][t]
			for leftBound := t; leftBound >= 0; leftBound-- {
				curMinHeight = min(curMinHeight, maxRectangleHeight[i][leftBound])
				countOfRectangle += curMinHeight
			}
		}
	}
	return countOfRectangle
}

func getRowsAndCols(matrix [][]int) (int, int) {
	if len(matrix) == 0 {
		return 0, 0
	}
	return len(matrix), len(matrix[0])
}

func get2DSlice(rows, column int) [][]int {
	slice := make([][]int, rows)
	for i := 0; i < len(slice); i++ {
		slice[i] = make([]int, column)
	}
	return slice
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// -------------------------- 单调栈 + DP --------------------------
// 时间复杂度: O(n * m)
func numSubmat(mat [][]int) int {
	maxRectangleHeight := getMaxRectangleHeight(mat)
	return getCountOfRectangle(maxRectangleHeight)
}

func getMaxRectangleHeight(mat [][]int) [][]int {
	rows, cols := getRowsAndCols(mat)
	maxRectangleHeight := get2DSlice(rows, cols)
	maxRectangleHeight[0][0] = mat[0][0]
	for i := 1; i < rows; i++ {
		if mat[i][0] == 1 {
			maxRectangleHeight[i][0] = maxRectangleHeight[i-1][0] + 1
		}
	}
	for t := 1; t < cols; t++ {
		maxRectangleHeight[0][t] = mat[0][t]
	}
	for i := 1; i < rows; i++ {
		for t := 1; t < cols; t++ {
			if mat[i][t] == 1 {
				maxRectangleHeight[i][t] = maxRectangleHeight[i-1][t] + 1
			}
		}
	}
	return maxRectangleHeight
}

func getCountOfRectangle(maxRectangleHeight [][]int) int {
	rows, cols := getRowsAndCols(maxRectangleHeight)
	countOfRectangle := 0
	for i := 0; i < rows; i++ {
		countOfRectangleInCurRow := make([]int, cols)
		stack := NewMyStack()
		for t := 0; t < cols; t++ {
			for !stack.IsEmpty() && maxRectangleHeight[i][stack.GetTop()] >= maxRectangleHeight[i][t] {
				stack.Pop()
			}
			if stack.IsEmpty() {
				rectangleWidth := t + 1
				countOfRectangleInCurRow[t] = maxRectangleHeight[i][t] * rectangleWidth
			} else {
				rectangleWidth := t - stack.GetTop()
				countOfRectangleInCurRow[t] = maxRectangleHeight[i][t]*rectangleWidth + countOfRectangleInCurRow[stack.GetTop()]
			}
			stack.Push(t)
			countOfRectangle += countOfRectangleInCurRow[t]
		}
	}
	return countOfRectangle
}

func getRowsAndCols(matrix [][]int) (int, int) {
	if len(matrix) == 0 {
		return 0, 0
	}
	return len(matrix), len(matrix[0])
}

func get2DSlice(rows, column int) [][]int {
	slice := make([][]int, rows)
	for i := 0; i < len(slice); i++ {
		slice[i] = make([]int, column)
	}
	return slice
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// -------------------------- MyStack --------------------------
type MyStack struct {
	data []int
}

func NewMyStack() *MyStack {
	return &MyStack{}
}

func (ms *MyStack) Push(val int) {
	ms.data = append(ms.data, val)
}

func (ms *MyStack) Pop() int {
	top := ms.data[ms.GetSize()-1]
	ms.data = ms.data[:ms.GetSize()-1]
	return top
}

func (ms *MyStack) GetTop() int {
	return ms.data[ms.GetSize()-1]
}

func (ms *MyStack) IsEmpty() bool {
	return ms.GetSize() == 0
}

func (ms *MyStack) GetSize() int {
	return len(ms.data)
}

/*
	总结:
		1. 单调栈 + DP 真的NB。
*/
