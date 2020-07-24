package 模拟

func isValid(S string) bool {
	stack := NewMyStack()
	for i := 0; i < len(S); i++ {
		switch {
		case S[i] == 'a':
			stack.Push(S[i])
		case S[i] == 'b':
			if stack.IsEmpty() || stack.GetTop() != 'a' {
				return false
			}
			stack.Push(S[i])
		case S[i] == 'c':
			if stack.IsEmpty() || stack.GetTop() != 'b' {
				return false
			}
			stack.Pop()
			stack.Pop()
		}
	}
	return stack.IsEmpty()
}

// ------------------------- MyStack -------------------------
type MyStack struct {
	data []byte
}

func NewMyStack() *MyStack {
	return &MyStack{}
}

func (ms *MyStack) Push(val byte) {
	ms.data = append(ms.data, val)
}

func (ms *MyStack) Pop() byte {
	top := ms.data[ms.GetSize()-1]
	ms.data = ms.data[:ms.GetSize()-1]
	return top
}

func (ms *MyStack) GetTop() byte {
	return ms.data[ms.GetSize()-1]
}

func (ms *MyStack) IsEmpty() bool {
	return ms.GetSize() == 0
}

func (ms *MyStack) GetSize() int {
	return len(ms.data)
}

/*
	题目链接: https://leetcode-cn.com/problems/check-if-word-is-valid-after-substitutions/
	总结;
		1. 我还记得有一道题是: 判断这个字符串是否只由 abc组成，abc之间可以穿插。
*/
