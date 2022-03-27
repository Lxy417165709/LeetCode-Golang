package struct_util

// MyStack 栈。
type MyStack struct {
	Elements []interface{} // 存放栈元素。
}

// NewMyStack 构造函数。
func NewMyStack() *MyStack {
	return &MyStack{}
}

// Push 入栈。
func (m *MyStack) Push(element interface{}) {
	m.Elements = append(m.Elements, element)
}

// Pop 出栈。
func (m *MyStack) Pop() interface{} {
	element := m.Elements[len(m.Elements)-1]
	m.Elements = m.Elements[:len(m.Elements)-1]
	return element
}

// Top 取顶。
func (m *MyStack) Top() interface{} {
	element := m.Elements[len(m.Elements)-1]
	return element
}

// IsEmpty 判空。
func (m *MyStack) IsEmpty() bool {
	return len(m.Elements) == 0
}
