package 链表

// --------------------- CQueue ---------------------
// 双栈。
// 		DeleteHead   时间复杂度为 O(n)
//		AppendTail  时间复杂度为 O(1)
type CQueue struct {
	pushStack *MyStack
	popStack *MyStack
}

func Constructor() CQueue {
	return CQueue{
		pushStack:NewMyStack(),
		popStack:NewMyStack(),
	}
}

func (mq *CQueue) AppendTail(value int)  {
	mq.pushStack.Push(value)
}

func (mq *CQueue) DeleteHead() int {
	if mq.popStack.IsEmpty(){
		mq.moveAllPushStackNumsToPopStack()
	}
	if mq.popStack.IsEmpty() {
		return -1
	}
	return mq.popStack.Pop()
}

func (mq *CQueue)moveAllPushStackNumsToPopStack() {
	for !mq.pushStack.IsEmpty() {
		mq.popStack.Push(mq.pushStack.Pop())
	}
}

// --------------------- CQueue ---------------------
// 双栈。
// 		DeleteHead   时间复杂度为 O(1)
//		AppendTail  时间复杂度为 O(n)
type CQueue struct {
	dataStack *MyStack
	auxiliaryStack *MyStack
}


func Constructor() CQueue {
	return CQueue{
		dataStack:NewMyStack(),
		auxiliaryStack:NewMyStack(),
	}
}

func (mq *CQueue) AppendTail(value int)  {
	mq.moveAllDataStackNumsToAuxiliaryStack()
	mq.dataStack.Push(value)
	mq.moveAllAuxiliaryStackNumsToDataStack()
}

func (mq *CQueue) DeleteHead() int {
	if mq.dataStack.IsEmpty(){
		return -1
	}
	return mq.dataStack.Pop()
}

func (mq *CQueue) moveAllDataStackNumsToAuxiliaryStack() {
	transferData(mq.dataStack,mq.auxiliaryStack)
}

func (mq *CQueue)moveAllAuxiliaryStackNumsToDataStack() {
	transferData(mq.auxiliaryStack,mq.dataStack)
}

func transferData(src,des *MyStack) {
	for !src.IsEmpty() {
		des.Push(src.Pop())
	}
}

// --------------------- MyStack ---------------------
type MyStack struct{
	data []int
}

func NewMyStack() *MyStack{
	return &MyStack{}
}

func (ms *MyStack) Push(x int) {
	ms.data = append(ms.data,x)
}

func (ms *MyStack) Pop() int{
	top := ms.data[ms.GetSize()-1]
	ms.data = ms.data[:ms.GetSize()-1]
	return top
}

func (ms *MyStack) Top() int{
	return ms.data[ms.GetSize()-1]
}

func (ms *MyStack) IsEmpty() bool{
	return ms.GetSize()==0
}

func (ms *MyStack) GetSize() int{
	return len(ms.data)
}

/*
	题目链接: https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/
*/
