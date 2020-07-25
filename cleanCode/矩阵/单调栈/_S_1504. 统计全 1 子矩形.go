package 单调栈

// -------------------------- 无单调栈，朴素解法 --------------------------
// 时间复杂度: O(n^2 * m^2)
func numSubmat(mat [][]int) int {
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
	result := 0
	for i := 0; i < rows; i++ {
		for t := 0; t < cols; t++ {
			for h := 1; h <= maxRectangleHeight[i][t]; h++ {
				// 纵向扩展
				result++

				// 横向扩展
				for leftBound := t - 1; leftBound >= 0 && maxRectangleHeight[i][leftBound] >= h; leftBound-- {
					result ++
				}

			}
		}
	}
	return result
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
	result := 0
	for i := 0; i < rows; i++ {
		for t := 0; t < cols; t++ {
			// 这里初现单调栈的端倪
			curMinHeight := maxRectangleHeight[i][t]
			for leftBound := t; leftBound >= 0; leftBound-- {
				curMinHeight = min(curMinHeight,maxRectangleHeight[i][leftBound])
				result += curMinHeight
			}
		}
	}
	return result
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

func min(a,b int) int{
	if a>b{
		return b
	}
	return a
}

// -------------------------- 单调栈 --------------------------
// 时间复杂度: O(n * m)
func numSubmat(mat [][]int) int {
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
	result := 0
	for i := 0; i < rows; i++ {
		result = max(result,maxRectangleHeight[i])
	}
	return result
}

// 套用 _84. 柱状图中最大的矩形_ 的代码
func largestRectangleArea(heights []int) int {
	indexStack := NewMyStack()
	largestArea:=0
	heights = append(heights,0) // 加入0的目的是: 为了让最终的单调递增栈被清空。 (清空栈的过程中会求取矩形的面积)
	for i:=0;i<len(heights);i++{
		for !indexStack.IsEmpty() && heights[indexStack.GetTop()]>=heights[i] {
			heightOfRectangle, widthOfRectangle := heights[indexStack.Pop()],0
			if indexStack.IsEmpty(){
				widthOfRectangle = i
			}else{
				widthOfRectangle = i - indexStack.GetTop() - 1
			}
			largestArea = max(largestArea,heightOfRectangle*widthOfRectangle)
		}
		indexStack.Push(i)
	}
	return largestArea
}

func max(a,b int) int{
	if a>b{
		return a
	}
	return b
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
