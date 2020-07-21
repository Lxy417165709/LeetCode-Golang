package 单调栈

func finalPrices(prices []int) []int {
	indexStack := NewMyStack()
	result := make([]int, len(prices))
	for i := 0; i < len(prices); i++ {
		for !indexStack.IsEmpty() && prices[indexStack.GetTop()] >= prices[i] {
			result[indexStack.GetTop()] = prices[indexStack.GetTop()] - prices[i]
			indexStack.Pop()
		}
		indexStack.Push(i)
	}
	for !indexStack.IsEmpty() {
		result[indexStack.GetTop()] = prices[indexStack.GetTop()]
		indexStack.Pop()
	}
	return result
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
	题目链接:
	总结:
		1. 这题就是找 A[i+1:] 中第一个大于等于 A[i]的数
*/
