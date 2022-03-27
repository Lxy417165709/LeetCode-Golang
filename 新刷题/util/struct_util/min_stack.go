package struct_util

// 注意: 这个代码还没测试。

// MySortedStack 排序栈。
type MySortedStack struct {
	stack        *MyStack
	compareValue func(interface{}, interface{}) bool // 栈内保证: compareValue(栈顶元素, 非栈顶元素) 必为 true。
}

func NewMyMinStack(compare func(interface{}, interface{}) bool) *MySortedStack {
	return &MySortedStack{stack: NewMyStack(), compareValue: compare}
}

func (m *MySortedStack) Push(val interface{}) {
	for !m.stack.IsEmpty() && m.compareValue(val, m.stack.Top()) {
		m.stack.Pop()
	}
	m.stack.Push(val)
}

func (m *MySortedStack) Pop() interface{} {
	return m.stack.Pop()
}

func (m *MySortedStack) Top() interface{} {
	return m.stack.Top()
}
