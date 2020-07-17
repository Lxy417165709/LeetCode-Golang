package 单调栈

// --------------------- 单调递增栈 ----------------------
// 执行用时：4 ms,   在所有 Go 提交中击败了 99.91% 的用户
// 内存消耗：5.8 MB, 在所有 Go 提交中击败了 33.33% 的用户
//
// 时间复杂度 O(n)
func largestRectangleArea(heights []int) int {
	indexStack := NewMyStack()
	largestArea := 0
	heights = append(heights, 0) // 加入 0 的目的是: 为了让最终的单调递增栈被清空。 (清空栈的过程中会求取矩形的面积)
	for i := 0; i < len(heights); i++ {
		for !indexStack.IsEmpty() && heights[indexStack.GetTop()] >= heights[i] {
			heightOfRectangle, widthOfRectangle := heights[indexStack.Pop()], 0
			if indexStack.IsEmpty() {
				widthOfRectangle = i // 即宽度为 区间[0, i) 的长度
			} else {
				widthOfRectangle = i - indexStack.GetTop() - 1 // 即宽度为 区间 (indexStack.GetTop(), i) 的长度
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

// --------------------- MyStack ---------------------
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
	题目链接: https://leetcode-cn.com/problems/largest-rectangle-in-histogram/
	总结:
		1. 这题就是求取: 能完全包含数组中柱子的矩形的面积最大值。 (也可以采用左右扩展的方式求得)
*/
