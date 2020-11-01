package 模拟

// ------------------- 方法1: 使用栈 -------------------
func minAddToMakeValid(S string) int {
	stackStoringLeftParentheses := NewMyStack()
	minCountOfAdding := 0
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			stackStoringLeftParentheses.Push(S[i])
			continue
		}
		if stackStoringLeftParentheses.IsEmpty() {
			minCountOfAdding++
			continue
		}
		stackStoringLeftParentheses.Pop()
	}
	minCountOfAdding += stackStoringLeftParentheses.GetSize()
	return minCountOfAdding
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

// ------------------- 方法2: 计次 -------------------
func minAddToMakeValid(S string) int {
	countOfLeftParentheses := 0
	minCountOfAdding := 0
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			countOfLeftParentheses++
			continue
		}
		if countOfLeftParentheses == 0 {
			minCountOfAdding++
			continue
		}
		countOfLeftParentheses--
	}
	minCountOfAdding += countOfLeftParentheses
	return minCountOfAdding
}

/*
	题目链接: https://leetcode-cn.com/problems/minimum-add-to-make-parentheses-valid/comments/
	总结:
		1. 其实栈是可以不要的，只需要记录当前拥有多少左括号就可以了。
*/
