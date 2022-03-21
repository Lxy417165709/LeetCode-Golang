package 链表


// --------------------- MyQueue ---------------------
// 双栈。
// 		Pop   时间复杂度为 O(n)
//		Insert  时间复杂度为 O(1)
type MyQueue struct {
	pushStack *MyStack
	popStack *MyStack
}

func Constructor() MyQueue {
	return MyQueue{
		pushStack:NewMyStack(),
		popStack:NewMyStack(),
	}
}

func (mq *MyQueue) Push(x int)  {
	mq.pushStack.Push(x)
}

func (mq *MyQueue) Pop() int {
	if mq.popStack.IsEmpty(){
		mq.moveAllPushStackNumsToPopStack()
	}
	return mq.popStack.Pop()
}

func (mq *MyQueue) Peek() int {
	if mq.popStack.IsEmpty(){
		mq.moveAllPushStackNumsToPopStack()
	}
	return mq.popStack.Top()
}

func (mq *MyQueue) Empty() bool {
	return mq.pushStack.IsEmpty() && mq.popStack.IsEmpty()
}

func (mq *MyQueue) moveAllPushStackNumsToPopStack() {
	for !mq.pushStack.IsEmpty() {
		mq.popStack.Push(mq.pushStack.Pop())
	}
}

// --------------------- MyQueue ---------------------
// 双栈。
// 		Pop   时间复杂度为 O(1)
//		Insert  时间复杂度为 O(n)
type MyQueue struct {
	dataStack *MyStack
	auxiliaryStack *MyStack
}

func Constructor() MyQueue {
	return MyQueue{
		dataStack:NewMyStack(),
		auxiliaryStack:NewMyStack(),
	}
}

func (mq *MyQueue) Push(x int)  {
	mq.moveAllDataStackNumsToAuxiliaryStack()
	mq.dataStack.Push(x)
	mq.moveAllAuxiliaryStackNumsToDataStack()
}

func (mq *MyQueue) Pop() int {
	return mq.dataStack.Pop()
}

func (mq *MyQueue) Peek() int {
	return mq.dataStack.Top()
}

func (mq *MyQueue) Empty() bool {
	return mq.dataStack.IsEmpty()
}

func (mq *MyQueue) moveAllDataStackNumsToAuxiliaryStack() {
	transferData(mq.dataStack,mq.auxiliaryStack)
}

func (mq *MyQueue) moveAllAuxiliaryStackNumsToDataStack() {
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
	题目链接: https://leetcode-cn.com/problems/implement-queue-using-stacks/
*/
