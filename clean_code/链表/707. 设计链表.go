package 链表

// --------------------------- MyLinkedList ---------------------------
// 解题方式: 双链表 + 哑头哑尾。
type MyLinkedList struct {
	dummyHead, dummyTail *Node
}

func Constructor() MyLinkedList {
	dummyHead, dummyTail := &Node{}, &Node{}
	dummyHead.Next, dummyTail.Pre = dummyTail, dummyHead
	return MyLinkedList{dummyHead: dummyHead, dummyTail: dummyTail}
}

func (ml *MyLinkedList) Get(index int) int {
	node := ml.getNthNode(index)
	if ml.isNodeValid(node) {
		return node.Val
	}
	return -1
}

func (ml *MyLinkedList) AddAtHead(val int) {
	linkAfterNodeA(ml.dummyHead, &Node{Val: val})
}

func (ml *MyLinkedList) AddAtTail(val int) {
	linkAfterNodeA(ml.dummyTail.Pre, &Node{Val: val})
}

func (ml *MyLinkedList) AddAtIndex(index int, val int) {
	node := ml.getNthNode(index - 1)
	if node == ml.dummyTail {
		return
	}
	linkAfterNodeA(node, &Node{Val: val})
}

func (ml *MyLinkedList) DeleteAtIndex(index int) {
	node := ml.getNthNode(index)
	if ml.isNodeValid(node) {
		eraseSelf(node)
	}
}

func (ml *MyLinkedList) getNthNode(index int) *Node {
	// index < 0，		  	  返回 this.dummyHead
	// 0 <= index < 链表长度, 返回 链表节点
	// index >= 链表长度，	  返回 this.dummyTail

	cur := ml.dummyHead
	for i := 0; i <= index && cur != ml.dummyTail; i++ {
		cur = cur.Next
	}
	return cur
}

func (ml *MyLinkedList) isNodeValid(nodeInList *Node) bool {
	return nodeInList != ml.dummyHead && nodeInList != ml.dummyTail
}

// --------------------------- Node ---------------------------
type Node struct {
	Val       int
	Next, Pre *Node
}

func linkAfterNodeA(nodeA, linkedNode *Node) {
	nodeA.Next.Pre = linkedNode
	linkedNode.Next = nodeA.Next
	nodeA.Next = linkedNode
	linkedNode.Pre = nodeA
}

func eraseSelf(node *Node) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}
