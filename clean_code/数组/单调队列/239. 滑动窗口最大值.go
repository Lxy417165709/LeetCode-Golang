package 单调队列

func maxSlidingWindow(nums []int, k int) []int {
	queue := NewMyQueue()
	result := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		for !queue.IsEmpty() && i-queue.GetTop() >= k {
			queue.Pop()
		}
		for !queue.IsEmpty() && nums[queue.GetTail()] < nums[i] {
			queue.PopTail()
		}
		queue.Push(i)
		if i >= k-1 {
			result = append(result, nums[queue.GetTop()])
		}
	}
	return result
}

// ------------------- MyQueue -------------------
type MyQueue struct {
	data []int
}

func NewMyQueue() *MyQueue {
	return &MyQueue{}
}

func (mq *MyQueue) IsEmpty() bool {
	return mq.GetSize() == 0
}

func (mq *MyQueue) Push(num int) {
	mq.data = append(mq.data, num)
}

func (mq *MyQueue) Pop() int {
	top := mq.GetTop()
	mq.data = mq.data[1:]
	return top
}

func (mq *MyQueue) PopTail() int {
	tail := mq.GetTail()
	mq.data = mq.data[:mq.GetSize()-1]
	return tail
}

func (mq *MyQueue) GetTop() int {
	return mq.data[0]
}

func (mq *MyQueue) GetTail() int {
	return mq.data[mq.GetSize()-1]
}

func (mq *MyQueue) GetSize() int {
	return len(mq.data)
}

/*
	题目链接: https://leetcode-cn.com/problems/sliding-window-maximum/
	总结:
		1. 单调队列与单调栈的区别在于: 单调队列的头可访问、移除。
*/
