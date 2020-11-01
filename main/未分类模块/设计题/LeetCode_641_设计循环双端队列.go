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

// ----------- MyCircularDeque -----------
type MyCircularDeque struct {
	dummyHead, dummyTail *Node
	length               int
	capacity             int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(capacity int) MyCircularDeque {
	dummyHead, dummyTail := NewNode(0), NewNode(0)
	dummyHead.Next = dummyTail
	dummyTail.Pre = dummyHead
	return MyCircularDeque{dummyHead, dummyTail, 0, capacity}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (mcd *MyCircularDeque) InsertFront(value int) bool {
	if mcd.length == mcd.capacity {
		return false
	}
	newNode := NewNode(value)
	newNode.Pre = mcd.dummyHead
	newNode.Next = mcd.dummyHead.Next
	mcd.dummyHead.Next.Pre = newNode
	mcd.dummyHead.Next = newNode
	mcd.length++
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (mcd *MyCircularDeque) InsertLast(value int) bool {
	if mcd.length == mcd.capacity {
		return false
	}
	newNode := NewNode(value)
	newNode.Next = mcd.dummyTail
	newNode.Pre = mcd.dummyTail.Pre
	mcd.dummyTail.Pre.Next = newNode
	mcd.dummyTail.Pre = newNode
	mcd.length++
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (mcd *MyCircularDeque) DeleteFront() bool {
	if mcd.length == 0 {
		return false
	}
	shouldDeleteNode := mcd.dummyHead.Next
	deleteNodeInList(shouldDeleteNode)
	mcd.length--
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (mcd *MyCircularDeque) DeleteLast() bool {
	if mcd.length == 0 {
		return false
	}
	shouldDeleteNode := mcd.dummyTail.Pre
	deleteNodeInList(shouldDeleteNode)
	mcd.length--
	return true
}

/** Get the front item from the deque. */
func (mcd *MyCircularDeque) GetFront() int {
	if mcd.length == 0 {
		return -1
	}
	return mcd.dummyHead.Next.Val
}

/** Get the last item from the deque. */
func (mcd *MyCircularDeque) GetRear() int {
	if mcd.length == 0 {
		return -1
	}
	return mcd.dummyTail.Pre.Val
}

/** Checks whether the circular deque is empty or not. */
func (mcd *MyCircularDeque) IsEmpty() bool {
	return mcd.length == 0
}

/** Checks whether the circular deque is full or not. */
func (mcd *MyCircularDeque) IsFull() bool {
	return mcd.length == mcd.capacity
}

func deleteNodeInList(node *Node) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
	node.Next = nil
	node.Pre = nil
}
