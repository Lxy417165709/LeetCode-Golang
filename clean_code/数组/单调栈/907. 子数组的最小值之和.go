package 单调栈

import "fmt"

// ---------------------------- 方法1: 纯暴力法 ----------------------------
// 1. 枚举区间的两个端点
// 2. 计算区间的最小值
// 3. 将这些最小值加起来就是答案了
//    时间复杂度 O(n^3)

// ---------------------------- 方法2: 优化暴力法 ----------------------------
// 执行用时：896 ms, 在所有 Go 提交中击败了 6.45%   的用户
// 内存消耗：6.5 MB, 在所有 Go 提交中击败了 100.00% 的用户
//    时间复杂度 O(n^2)

const mod = 1000000007

func sumSubarrayMins(A []int) int {
	ans := 0
	for i := 0; i < len(A); i++ {
		leftCount := getLeftSideCountOfGreater(A, i)
		rightCount := getRightSideCountOfGreaterOrEqual(A, i)
		ans += (leftCount + 1) * (rightCount + 1) * A[i]
		ans %= mod
	}
	return ans
}

func getLeftSideCountOfGreater(A []int, indexOfNum int) int {
	count := 0
	curIndex := indexOfNum - 1
	for curIndex >= 0 && A[curIndex] > A[indexOfNum] {
		curIndex--
		count++
	}
	return count
}

func getRightSideCountOfGreaterOrEqual(A []int, indexOfNum int) int {
	count := 0
	curIndex := indexOfNum + 1
	for curIndex < len(A) && A[curIndex] >= A[indexOfNum] {
		curIndex++
		count++
	}
	return count
}

// ---------------------------- 方法3: 利用单调递增栈优化方法2 ----------------------------
// 执行用时：68 ms,  在所有 Go 提交中击败了 67.74% 的用户
// 内存消耗：6.7 MB, 在所有 Go 提交中击败了 50.00% 的用户
//    时间复杂度 O(n)
const mod = 1000000007

func sumSubarrayMins(A []int) int {
	leftCount := getRelationOfIndexToLeftSideCountOfGreater(A)
	rightCount := getRelationOfIndexToRightSideCountOfGreaterOrEqual(A)
	ans := 0
	for i := 0; i < len(A); i++ {
		ans += (leftCount[i] + 1) * (rightCount[i] + 1) * A[i]
		ans %= mod
	}
	return ans
}

func getRelationOfIndexToLeftSideCountOfGreater(A []int) []int {
	stack := NewMyStack()
	leftSideCountOfGreater := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		for !stack.IsEmpty() && A[i] <= A[stack.GetTop()] {
			countOfIntervalNums := i - stack.GetTop() // 即 [stack.GetTop(), i) 区间内一共有多少个数字
			leftSideCountOfGreater[i] = countOfIntervalNums + leftSideCountOfGreater[stack.GetTop()]
			stack.Pop()
		}
		stack.Push(i)
	}
	return leftSideCountOfGreater
}

func getRelationOfIndexToRightSideCountOfGreaterOrEqual(A []int) []int {
	stack := NewMyStack()
	rightSideCountOfGreaterOrEqual := make([]int, len(A))
	for i := len(A) - 1; i >= 0; i-- {
		for !stack.IsEmpty() && A[i] <= A[stack.GetTop()] {
			countOfIntervalNums := stack.GetTop() - i // 即 (i, stack.GetTop()] 区间内一共有多少个数字
			rightSideCountOfGreaterOrEqual[i] = countOfIntervalNums + rightSideCountOfGreaterOrEqual[stack.GetTop()]
			stack.Pop()
		}
		stack.Push(i)
	}
	return rightSideCountOfGreaterOrEqual
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

// ---------------------------- 方法4: 单调递增栈 + DP ----------------------------
// 执行用时：68 ms,  在所有 Go 提交中击败了 67.74%  的用户
// 内存消耗：6.7 MB, 在所有 Go 提交中击败了 100.00% 的用户
//    时间复杂度 O(n)

const mod = 1000000007

func sumSubarrayMins(A []int) int {
	minSumBackByIndex := make([]int, len(A)) // 定义: 所有以 A[i] 结尾的子序列的最小和，
	// 如 [100 200 300]
	// 则 minSumBackByIndex[2] 表示: [300], [200 300], [100 200 300] 的最小值之和，
	// 于是 minSumBackByIndex[2] = 300 + 200 + 100 = 600

	stack := NewMyStack()
	for i := 0; i < len(A); i++ {
		for !stack.IsEmpty() && A[i] <= A[stack.GetTop()] {
			stack.Pop()
		}
		if stack.IsEmpty() {
			// 表示 A[i] 大于等于 A[:i] 中的所有元素
			minSumBackByIndex[i] = A[i] * (i + 1)
		} else {
			// 表示 A[i] 大于等于 A[topIndex+1:i] 中的所有元素
			topIndex := stack.GetTop()
			minSumBackByIndex[i] = A[i]*(i-topIndex) + minSumBackByIndex[topIndex]
		}
		minSumBackByIndex[i] %= mod
		stack.Push(i)
	}
	return getSumModM(minSumBackByIndex, mod)
}

func getSumModM(array []int, m int) int {
	sum := 0
	for i := 0; i < len(array); i++ {
		sum += array[i]
		sum %= m
	}
	return sum
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
	总结:
		1. 注意，左边求的是大于，右边是大于等于，这样就不会出现重复，如果都是大于等于，那这会出现重复。
				(左边大于等于，右边大于也可以)
		2. 进阶做法是维护单调递增栈。
		3. 方法4的dp也用到了单调栈。
*/
