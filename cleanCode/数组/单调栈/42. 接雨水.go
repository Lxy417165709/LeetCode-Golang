package 单调栈

import "fmt"

// -------------------------- 方法1: 记录左右两边最高的柱子 --------------------------
// 执行用时：4 ms,   在所有 Go 提交中击败了 88.45% 的用户
// 内存消耗：3.1 MB, 在所有 Go 提交中击败了 33.33% 的用户
//
// 方法1 基于垂直方向接雨水，每一次只求指定柱子能接收到多少雨水。
// 时间复杂度 O(n)

func trap(height []int) int {
	return getReceivedRain(height)
}

func getReceivedRain(height []int) int {
	if len(height) == 0 {
		return 0
	}
	leftPartHighestBar := getLeftPartHighestBar(height)
	rightPartHighestBar := getRightPartHighestBar(height)
	receivedRain := 0
	for i := 1; i <= len(height)-2; i++ {
		rain := max(min(leftPartHighestBar[i-1], rightPartHighestBar[i+1])-height[i], 0)
		receivedRain += rain
	}
	return receivedRain
}

func getLeftPartHighestBar(height []int) []int {
	leftPartHighestBar := make([]int, len(height))
	leftPartHighestBar[0] = height[0]
	for i := 1; i < len(height); i++ {
		leftPartHighestBar[i] = max(leftPartHighestBar[i-1], height[i])
	}
	return leftPartHighestBar
}

func getRightPartHighestBar(height []int) []int {
	rightPartHighestBar := make([]int, len(height))
	rightPartHighestBar[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rightPartHighestBar[i] = max(rightPartHighestBar[i+1], height[i])
	}
	return rightPartHighestBar
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// -------------------------- 方法2: 维护一个单调递减栈 --------------------------
// 执行用时：4 ms,   在所有 Go 提交中击败了 88.45% 的用户
// 内存消耗：2.8 MB, 在所有 Go 提交中击败了 66.67% 的用户
//
// 方法2 基于水平方向接雨水，每一次求每个洼地能接收多少雨水。
// 什么是洼地？就是能接收雨水的地方，如 [5 1 1 7 5], 索引0 到 索引1 就算是个洼地。
// 时间复杂度 O(n)

func trap(height []int) int {
	return getReceivedRain(height)
}
func getReceivedRain(height []int) int {
	stack := NewMyStack()
	receivedRain := 0
	for i := 0; i < len(height); i++ {
		for !stack.IsEmpty() && height[stack.GetTop()] <= height[i] {
			bottomHeightOfDepression := height[stack.Pop()] // 洼地的高度
			if stack.IsEmpty() {
				break
			}
			widthOfDepression := i - stack.GetTop() - 1                              // 对应洼地的宽度
			minBoundOfDepression := min(height[i], height[stack.GetTop()])           // 对应洼地最小左右边界的高度。
			heightOfReceivingRain := minBoundOfDepression - bottomHeightOfDepression // 对应该洼地能接收雨水的高度
			receivedRain += widthOfDepression * heightOfReceivingRain
		}
		stack.Push(i)
	}
	return receivedRain
}

// ------------------------- MyStack -------------------------
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*
	题目链接: https://leetcode-cn.com/problems/trapping-rain-water/submissions/
	总结:
		1. 单调栈好玩~
*/
