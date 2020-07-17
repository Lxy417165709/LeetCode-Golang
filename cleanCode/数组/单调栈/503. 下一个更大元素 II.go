package 单调栈

// ----------------------------- 方法1: 单调递减栈 -----------------------------
// 执行用时：32 ms,  在所有 Go 提交中击败了 90.26%  的用户
// 内存消耗：6.8 MB, 在所有 Go 提交中击败了 100.00% 的用户
//
// 时间复杂度 O(n)

func nextGreaterElements(nums []int) []int {
	result := getNextGreaterNums(append(nums, nums...))
	return reverse(result[len(nums):])
}

func getNextGreaterNums(nums []int) []int {
	valStack := NewMyStack()
	nextGreaterNums := make([]int, 0)
	for i := len(nums) - 1; i >= 0; i-- {
		for !valStack.IsEmpty() && valStack.GetTop() <= nums[i] {
			valStack.Pop()
		}
		if !valStack.IsEmpty() {
			nextGreaterNums = append(nextGreaterNums, valStack.GetTop())
		} else {
			nextGreaterNums = append(nextGreaterNums, -1)
		}
		valStack.Push(nums[i])
	}
	return nextGreaterNums
}

func reverse(arr []int) []int {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	return arr
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
	题目链接: https://leetcode-cn.com/problems/next-greater-element-ii/submissions/
	总结:
		1. 涉及到循环的，一般都可以将原数组进行拼接，之后再求解。
*/
