package 栈与队列

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

// Push 出栈。
func (m *MyStack) Pop() interface{} {
	element := m.Elements[len(m.Elements)-1]
	m.Elements = m.Elements[:len(m.Elements)-1]
	return element
}

// IsEmpty 判空。
func (m *MyStack) IsEmpty() bool {
	return len(m.Elements) == 0
}
