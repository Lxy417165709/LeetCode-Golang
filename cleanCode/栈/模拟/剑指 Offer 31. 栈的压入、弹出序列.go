package 模拟

func validateStackSequences(pushed []int, popped []int) bool {
	stack := NewMyStack()
	curTopIndexInPoppedSequence := 0
	for _, num := range pushed {
		stack.Push(num)
		for !stack.IsEmpty() && stack.GetTop() == popped[curTopIndexInPoppedSequence] {
			stack.Pop()
			curTopIndexInPoppedSequence++
		}
	}
	return stack.IsEmpty()
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


/*
	题目链接: https://leetcode-cn.com/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof/
*/
