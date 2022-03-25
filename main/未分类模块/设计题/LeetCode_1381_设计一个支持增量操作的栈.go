package main

const (
	emptyFlag = -1
)

// ------------ Node ------------
type Node struct {
	Val       int
	Pre, Next *Node
}

func NewNode(val int) *Node {
	return &Node{val, nil, nil}
}

// erase the link between the node and its surrounding nodes,
// and keep the connectivity of the linkList
func (node *Node) BreakLink() {
	node.Next.Pre = node.Pre
	node.Pre.Next = node.Next
	node.Next = nil
	node.Pre = nil
}

// ------------ CustomStack ------------
type CustomStack struct {
	dummyHead, dummyTail *Node
	capacity             int
	length               int
}

func Constructor(maxSize int) CustomStack {
	dummyHead, dummyTail := NewNode(-1), NewNode(-1)
	dummyHead.Next = dummyTail
	dummyTail.Pre = dummyHead
	return CustomStack{dummyHead, dummyTail, maxSize, 0}
}
func (cs *CustomStack) Push(value int) {
	if cs.isFull() {
		return
	}
	cs.pushExec(value)
}
func (cs *CustomStack) Pop() int {
	if cs.isEmpty() {
		return emptyFlag
	}
	return cs.popExec()
}
func (cs *CustomStack) Increment(k int, val int) {
	cs.incrementExec(k, val)
}

func (cs *CustomStack) isFull() bool {
	return cs.length == cs.capacity
}
func (cs *CustomStack) isEmpty() bool {
	return cs.length == 0
}
func (cs *CustomStack) pushExec(value int) {
	newNode := NewNode(value)
	newNode.Next = cs.dummyTail
	newNode.Pre = cs.dummyTail.Pre
	cs.dummyTail.Pre.Next = newNode
	cs.dummyTail.Pre = newNode
	cs.length++
}
func (cs *CustomStack) popExec() int {
	popNode := cs.dummyTail.Pre
	popNode.BreakLink()
	cs.length--
	return popNode.Val
}
func (cs *CustomStack) incrementExec(k int, val int) {
	readingNode := cs.dummyHead
	readCount := 0
	for readCount < k && readingNode.Next != cs.dummyTail {
		readingNode.Next.Val += val
		readingNode = readingNode.Next
		readCount++
	}
}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
