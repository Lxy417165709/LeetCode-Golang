package main

// --------------------- StockSpanner ---------------------
// 执行用时：192 ms, 在所有 Go 提交中击败了 98.48%  的用户
// 内存消耗：8.7 MB, 在所有 Go 提交中击败了 100.00% 的用户
//
// 时间复杂度: O(n)
// 思路: 维护一个单调递减栈。

type StockSpanner struct {
	data       []int
	indexStack *MyStack
}

func Constructor() StockSpanner {
	return StockSpanner{
		data:       make([]int, 0),
		indexStack: NewMyStack(),
	}
}

func (ss *StockSpanner) Next(price int) int {
	ss.data = append(ss.data, price)
	curIndex := len(ss.data) - 1
	for !ss.indexStack.IsEmpty() && ss.data[ss.indexStack.GetTop()] <= ss.data[curIndex] {
		ss.indexStack.Pop()
	}
	maxCountOfContinuousDay := 0
	if ss.indexStack.IsEmpty() {
		maxCountOfContinuousDay = curIndex + 1
	} else {
		maxCountOfContinuousDay = curIndex - ss.indexStack.GetTop()
	}
	ss.indexStack.Push(curIndex)
	return maxCountOfContinuousDay
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
	题目链接:
	总结:
		1. 这题就是求: 数组中，A[:i+1] 中连续有多少个数小于等于 A[i]。
*/
