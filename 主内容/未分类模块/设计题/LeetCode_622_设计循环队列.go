package main
// ----------- Node -----------
type Node struct {
	Val       int
	Pre, Next *Node
}

func NewNode(val int) *Node {
	return &Node{val, nil, nil}
}
func (n *Node) GetVal() int {
	return n.Val
}

// ----------- MyCircularQueue -----------
type MyCircularQueue struct {
	dummyHead, dummyTail *Node
	length               int
	capacity             int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(capacity int) MyCircularQueue {
	dummyHead, dummyTail := NewNode(0), NewNode(0)
	dummyHead.Next = dummyTail
	dummyTail.Pre = dummyHead
	return MyCircularQueue{dummyHead, dummyTail, 0, capacity}
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (mcq *MyCircularQueue) EnQueue(value int) bool {
	if mcq.length == mcq.capacity {
		return false
	}
	newNode := NewNode(value)
	newNode.Next = mcq.dummyTail
	newNode.Pre = mcq.dummyTail.Pre
	mcq.dummyTail.Pre.Next = newNode
	mcq.dummyTail.Pre = newNode
	mcq.length++
	return true
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (mcq *MyCircularQueue) DeQueue() bool {
	if mcq.length == 0 {
		return false
	}
	shouldDeleteNode := mcq.dummyHead.Next
	deleteNodeInList(shouldDeleteNode)
	mcq.length--
	return true
}

/** Get the front item from the deque. */
func (mcq *MyCircularQueue) Front() int {
	if mcq.length == 0 {
		return -1
	}
	return mcq.dummyHead.Next.Val
}

/** Get the last item from the deque. */
func (mcq *MyCircularQueue) Rear() int {
	if mcq.length == 0 {
		return -1
	}
	return mcq.dummyTail.Pre.Val
}

/** Checks whether the circular deque is empty or not. */
func (mcq *MyCircularQueue) IsEmpty() bool {
	return mcq.length == 0
}

/** Checks whether the circular deque is full or not. */
func (mcq *MyCircularQueue) IsFull() bool {
	return mcq.length == mcq.capacity
}

func deleteNodeInList(node *Node) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
	node.Next = nil
	node.Pre = nil
}
