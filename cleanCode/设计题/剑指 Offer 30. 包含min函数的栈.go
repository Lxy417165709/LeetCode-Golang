package main


// ----------------------- 方法1: 双栈实现 -----------------------
const INF = 1000000000000

type MinStack struct {
	dataStack     *MyStack
	minValueStack *MyStack
}

func Constructor() MinStack {
	return MinStack{
		dataStack:     NewMyStack(),
		minValueStack: NewMyStack(),
	}
}

func (ms *MinStack) Push(x int) {
	nextMinValue := ms.minValueStack.GetMinValueBetweenTopAndReference(x)
	ms.dataStack.Push(x)
	ms.minValueStack.Push(nextMinValue)
}

func (ms *MinStack) Pop() {
	ms.dataStack.Pop()
	ms.minValueStack.Pop()
}

func (ms *MinStack) Top() int {
	return ms.dataStack.GetTop()
}

func (ms *MinStack) Min() int {
	return ms.minValueStack.GetTop()
}

// ----------------------- 方法2: 单栈实现 -----------------------
const INF = 1000000000000

type MinStack struct {
	compositeStack *MyStack
}

func Constructor() MinStack {
	return MinStack{
		compositeStack: NewMyStack(),
	}
}

func (ms *MinStack) Push(x int) {
	nextMinValue := ms.compositeStack.GetMinValueBetweenTopAndReference(x)

	ms.compositeStack.Push(x)
	ms.compositeStack.Push(nextMinValue)
}

func (ms *MinStack) Pop() {
	ms.compositeStack.Pop()
	ms.compositeStack.Pop()
}

func (ms *MinStack) Top() int {
	defer ms.compositeStack.Push(ms.compositeStack.Pop())
	return ms.compositeStack.GetTop()

	// 相当于
	//  minValue := ms.compositeStack.Pop()
	//	realValue := ms.compositeStack.GetTop()
	//	ms.compositeStack.Push(minValue)
	//	return realValue
}

func (ms *MinStack) Min() int {
	return ms.compositeStack.GetTop()
}

// ------------------- MyStack -------------------
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

func (ms *MyStack) GetMinValueBetweenTopAndReference(reference int) int{
	nextMinValue := INF
	if !ms.IsEmpty() {
		nextMinValue = min(nextMinValue, ms.GetTop())
	}
	return min(nextMinValue, reference)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*
	题目链接: https://leetcode-cn.com/problems/bao-han-minhan-shu-de-zhan-lcof/comments/
*/

