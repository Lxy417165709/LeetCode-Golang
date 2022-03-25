package 链表

// ------------------ MyStack ---------------------
// 双队列。
// 		Pop   时间复杂度为 O(n)
//		Push  时间复杂度为 O(1)
type MyStack struct {
	queueStoringNums   *MyQueue
	queueStoringTopNum *MyQueue // 这个队列最多只有1个元素，存放栈顶元素
}

func Constructor() MyStack {
	return MyStack{
		queueStoringNums:   NewMyQueue(),
		queueStoringTopNum: NewMyQueue(),
	}
}

func (ms *MyStack) Push(x int) {
	ms.moveTopNumToQueueStoringNums()
	ms.queueStoringNums.Push(x)
}

func (ms *MyStack) Pop() int {
	if ms.queueStoringTopNum.IsEmpty() {
		ms.moveQueueStoringNumsToQueueStoringTopNumUntilLeavingOne()
		ms.exchangeQueuesResponsibility()
	}
	return ms.queueStoringTopNum.Pop()
}

func (ms *MyStack) Top() int {
	if ms.queueStoringTopNum.IsEmpty() {
		ms.moveQueueStoringNumsToQueueStoringTopNumUntilLeavingOne()
		ms.exchangeQueuesResponsibility()
	}
	return ms.queueStoringTopNum.Top()
}

func (ms *MyStack) Empty() bool {
	return ms.queueStoringNums.IsEmpty() && ms.queueStoringTopNum.IsEmpty()
}

func (ms *MyStack) moveQueueStoringNumsToQueueStoringTopNumUntilLeavingOne() {
	for ms.queueStoringNums.GetSize() != 1 {
		ms.queueStoringTopNum.Push(ms.queueStoringNums.Pop())
	}
}

func (ms *MyStack) moveTopNumToQueueStoringNums() {
	for !ms.queueStoringTopNum.IsEmpty() {
		ms.queueStoringNums.Push(ms.queueStoringTopNum.Pop())
	}
}

func (ms *MyStack) exchangeQueuesResponsibility() {
	ms.queueStoringNums, ms.queueStoringTopNum = ms.queueStoringTopNum, ms.queueStoringNums
}

// ------------------ MyStack ---------------------
// 双队列。
// 		Pop   时间复杂度为 O(1)
//		Push  时间复杂度为 O(n)
type MyStack struct {
	popQueue  *MyQueue
	pushQueue *MyQueue
}

func Constructor() MyStack {
	return MyStack{
		queueStoringNums:   NewMyQueue(),
		queueStoringTopNum: NewMyQueue(),
	}
}

func (ms *MyStack) Push(x int) {
	ms.pushQueue.Push(x)
	ms.moveAllPopQueueNumsToPushQueue()
	ms.exchangeQueuesResponsibility()
}

func (ms *MyStack) Pop() int {
	return ms.popQueue.Pop()
}

func (ms *MyStack) Top() int {
	return ms.popQueue.Top()
}

func (ms *MyStack) Empty() bool {
	return ms.popQueue.IsEmpty()
}

func (ms *MyStack) moveAllPopQueueNumsToPushQueue() {
	for !ms.popQueue.IsEmpty() {
		ms.pushQueue.Push(ms.popQueue.Pop())
	}
}

func (ms *MyStack) exchangeQueuesResponsibility() {
	ms.popQueue, ms.pushQueue = ms.pushQueue, ms.popQueue
}

// ------------------ MyStack ---------------------
// 单队列。
// 		Pop   时间复杂度为 O(1)
//		Push  时间复杂度为 O(n)
type MyStack struct {
	queue *MyQueue
}

func Constructor() MyStack {
	return MyStack{
		queue: NewMyQueue(),
	}
}

func (ms *MyStack) Push(x int) {
	sizeBeforePushing := ms.queue.GetSize()
	ms.queue.Push(x)
	for i := 0; i < sizeBeforePushing; i++ {
		ms.queue.Push(ms.queue.Pop())
	}
}

func (ms *MyStack) Pop() int {
	return ms.queue.Pop()
}

func (ms *MyStack) Top() int {
	return ms.queue.Top()
}

func (ms *MyStack) Empty() bool {
	return ms.queue.IsEmpty()
}

// ------------------ MyQueue ---------------------
type MyQueue struct {
	data []int
}

func NewMyQueue() *MyQueue {
	return &MyQueue{}
}

func (mq *MyQueue) Push(val int) {
	mq.data = append(mq.data, val)
}

func (mq *MyQueue) IsEmpty() bool {
	return mq.GetSize() == 0
}

func (mq *MyQueue) GetSize() int {
	return len(mq.data)
}

func (mq *MyQueue) Pop() int {
	top := mq.data[0]
	mq.data = mq.data[1:]
	return top
}

func (mq *MyQueue) Top() int {
	return mq.data[0]
}

/*
	题目链接:
	总结:
		1. 单队列实现栈的想法真的好。
*/
