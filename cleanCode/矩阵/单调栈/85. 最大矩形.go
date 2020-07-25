package 单调栈

func maximalRectangle(mat [][]byte) int {
	if len(mat) == 0 {
		return 0
	}
	rows, cols := getRowsAndCols(mat)
	maxRectangleHeight := get2DSlice(rows, cols)
	maxRectangleHeight[0][0] = int(mat[0][0] - '0')
	for i := 1; i < rows; i++ {
		if mat[i][0] == '1' {
			maxRectangleHeight[i][0] = maxRectangleHeight[i-1][0] + 1
		}
	}
	for t := 1; t < cols; t++ {
		maxRectangleHeight[0][t] = int(mat[0][t]) - '0'
	}
	for i := 1; i < rows; i++ {
		for t := 1; t < cols; t++ {
			if mat[i][t] == '1' {
				maxRectangleHeight[i][t] = maxRectangleHeight[i-1][t] + 1
			}
		}
	}
	result := 0
	for i := 0; i < rows; i++ {
		result = max(result, largestRectangleArea(maxRectangleHeight[i]))
	}
	return result
}

// 套用 _84. 柱状图中最大的矩形_ 的代码
func largestRectangleArea(heights []int) int {
	indexStack := NewMyStack()
	largestArea := 0
	heights = append(heights, 0) // 加入0的目的是: 为了让最终的单调递增栈被清空。 (清空栈的过程中会求取矩形的面积)
	for i := 0; i < len(heights); i++ {
		for !indexStack.IsEmpty() && heights[indexStack.GetTop()] >= heights[i] {
			heightOfRectangle, widthOfRectangle := heights[indexStack.Pop()], 0
			if indexStack.IsEmpty() {
				widthOfRectangle = i
			} else {
				widthOfRectangle = i - indexStack.GetTop() - 1
			}
			largestArea = max(largestArea, heightOfRectangle*widthOfRectangle)
		}
		indexStack.Push(i)
	}
	return largestArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getRowsAndCols(matrix [][]byte) (int, int) {
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

/*
	总结:
		1. 这题是 _84. 柱状图中最大的矩形_ 的进阶。
*/
