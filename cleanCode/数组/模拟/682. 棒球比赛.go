package 模拟

import "strconv"

func calPoints(ops []string) int {
	stack := NewMyStack()
	pointSum := 0
	for i := 0; i < len(ops); i++ {
		switch {
		case ops[i] == "+":
			goalOfLastRound := stack.Pop()
			goalOfLastTwoRound := stack.Pop()
			stack.Push(goalOfLastTwoRound)
			stack.Push(goalOfLastRound)
			stack.Push(goalOfLastRound + goalOfLastTwoRound)
		case ops[i] == "D":
			goalOfLastRound := stack.GetTop()
			stack.Push(2 * goalOfLastRound)
		case ops[i] == "C":
			stack.Pop()
		default:
			goal, _ := strconv.Atoi(ops[i])
			stack.Push(goal)
		}
	}
	for !stack.IsEmpty() {
		pointSum += stack.Pop()
	}
	return pointSum
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
	题目链接: https://leetcode-cn.com/problems/baseball-game/submissions/
*/
