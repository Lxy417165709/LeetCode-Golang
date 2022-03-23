package 混合

type CQueue struct {
	StackOfTopIsQueueFirstElement *MyStack // 栈，栈顶元素为队列首元素。
	StackOfTopIsQueueLastElement  *MyStack // 栈，栈顶元素为队列尾元素。
}

func Constructor() CQueue {
	return CQueue{
		StackOfTopIsQueueFirstElement: NewMyStack(),
		StackOfTopIsQueueLastElement:  NewMyStack(),
	}
}

func (q *CQueue) AppendTail(value int) {
	q.StackOfTopIsQueueLastElement.Push(value)
}

func (q *CQueue) DeleteHead() int {
	// 1. StackOfTopIsQueueFirstElement 为空，将 StackOfTopIsQueueLastElement 元素全部推入，栈顶元素即为队列的首元素。
	if q.StackOfTopIsQueueFirstElement.IsEmpty() {
		for !q.StackOfTopIsQueueLastElement.IsEmpty() {
			q.StackOfTopIsQueueFirstElement.Push(q.StackOfTopIsQueueLastElement.Pop())
		}
	}

	// 2. StackOfTopIsQueueFirstElement 为空，返回-1。
	if q.StackOfTopIsQueueFirstElement.IsEmpty() {
		return -1
	}

	// 3. StackOfTopIsQueueFirstElement 非空，返回。
	return q.StackOfTopIsQueueFirstElement.Pop().(int)
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
