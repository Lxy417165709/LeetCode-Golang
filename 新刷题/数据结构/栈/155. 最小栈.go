package 栈

import "github.com/Lxy417165709/LeetCode-Golang/新刷题/util/struct_util"

type MinStack struct {
	stack         *struct_util.MyStack // 栈内是所有元素。
	minValueStack *struct_util.MyStack // 栈内元素来自 stack。保留截止至当前，栈内的最小元素。。

	// 以上两个栈，长度相等。
}

func Constructor() MinStack {
	return MinStack{
		minValueStack: struct_util.NewMyStack(),
		stack:         struct_util.NewMyStack(),
	}
}

func (m *MinStack) Push(val int) {
	// 1. 保存原栈元素。
	m.stack.Push(val)
	if m.minValueStack.IsEmpty() {
		m.minValueStack.Push(val)
		return
	}

	// 2. 保留截止至当前，栈内的最小元素。
	m.minValueStack.Push(min(m.minValueStack.Top().(int), val))
}

func (m *MinStack) Pop() {
	m.stack.Pop()
	m.minValueStack.Pop()
}

func (m *MinStack) Top() int {
	return m.stack.Top().(int)
}

func (m *MinStack) GetMin() int {
	return m.minValueStack.Top().(int)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
