package 单调栈

// ----------------------------- 方法1: 暴力法 -----------------------------
// 执行用时：4 ms,   在所有 Go 提交中击败了 89.07%  的用户
// 内存消耗：2.9 MB, 在所有 Go 提交中击败了 100.00% 的用户
//
// 	时间复杂度 O( len(num1)*len(num2) )

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)
	for i := 0; i < len(nums1); i++ {
		result = append(result, getNextGreaterNum(nums2, nums1[i]))
	}
	return result
}

func getNextGreaterNum(nums []int, ref int) int {
	hasRefNumAppeared := false
	for i := 0; i < len(nums); i++ {
		if nums[i] == ref {
			hasRefNumAppeared = true
		}
		if hasRefNumAppeared && nums[i] > ref {
			return nums[i]
		}
	}
	return -1
}

// ----------------------------- 方法2: 从右到左遍历，维护一个单调递减栈 -----------------------------
// 执行用时：4 ms,   在所有 Go 提交中击败了 89.35% 的用户
// 内存消耗：3.5 MB, 在所有 Go 提交中击败了 100.00% 的用户
//
// 时间复杂度 O( len(nums1) + len(nums2) )
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	nextGreaterNum := getMapOfNextGreaterNum(nums2)
	result := make([]int, 0)
	for i := 0; i < len(nums1); i++ {
		result = append(result, nextGreaterNum[nums1[i]])
	}
	return result
}

func getMapOfNextGreaterNum(nums []int) map[int]int {
	valStack := NewMyStack()
	nextGreaterNum := make(map[int]int)
	for i := len(nums) - 1; i >= 0; i-- {
		for !valStack.IsEmpty() && valStack.GetTop() <= nums[i] {
			valStack.Pop()
		}
		if !valStack.IsEmpty() {
			nextGreaterNum[nums[i]] = valStack.GetTop()
		} else {
			nextGreaterNum[nums[i]] = -1
		}
		valStack.Push(nums[i])
	}
	return nextGreaterNum
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
	题目链接: https://leetcode-cn.com/problems/next-greater-element-i/submissions/
*/
