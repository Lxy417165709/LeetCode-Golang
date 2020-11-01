package 单调栈

// -------------------------- 维护一个单调递减栈 --------------------------
// 执行用时：64 ms,  在所有 Go 提交中击败了 73.78%  的用户
// 内存消耗：6.8 MB, 在所有 Go 提交中击败了 100.00% 的用户
//
// 时间复杂度: O(n)
func dailyTemperatures(T []int) []int {
	indexStack := NewMyStack()
	countOfLessWaitDay := make([]int, len(T))
	for i := len(T) - 1; i >= 0; i-- {
		for !indexStack.IsEmpty() && T[indexStack.GetTop()] <= T[i] {
			indexStack.Pop()
		}
		if indexStack.IsEmpty() {
			countOfLessWaitDay[i] = 0
		} else {
			countOfLessWaitDay[i] = indexStack.GetTop() - i
		}
		indexStack.Push(i)
	}
	return countOfLessWaitDay
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
	总结:
		1. 这题就是求:
			假如 存在 A[i+1:] 的索引x, 且 A[x] 是第一个大于 A[i]的元素，求这两个位置的距离，如果不存在 x，则距离为 0。

*/
