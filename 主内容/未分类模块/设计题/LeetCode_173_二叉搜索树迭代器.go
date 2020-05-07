package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	buffer []*TreeNode
}

func NewStack() *Stack {
	return &Stack{make([]*TreeNode, 0)}
}
func (stk *Stack) Push(node *TreeNode) {
	stk.buffer = append(stk.buffer, node)
}
func (stk *Stack) Pop() (lastNode *TreeNode) {
	if stk.isEmpty() {
		lastNode = nil
		return
	}
	lastNode = stk.getLastNode()
	stk.removeLastNode()
	return
}

func (stk *Stack) isEmpty() bool {
	return stk.length() == 0
}
func (stk *Stack) length() int {
	return len(stk.buffer)
}
func (stk *Stack) getLastNode() *TreeNode {
	return stk.buffer[stk.length()-1]
}
func (stk *Stack) removeLastNode() {
	stk.buffer= stk.buffer[:stk.length()-1]
}

type BSTIterator struct {
	stack *Stack
}

func Constructor(root *TreeNode) BSTIterator {
	bsi := BSTIterator{NewStack()}
	bsi.iterate(root)
	return bsi
}

/** @return the next smallest number */
func (bsi *BSTIterator) Next() int {
	// the job has pointed that Next() is always valid
	treeNodeOfSmallestNumber := bsi.getTreeNodeOfSmallestNumber()
	bsi.iterate(treeNodeOfSmallestNumber.Right)
	return treeNodeOfSmallestNumber.Val
}

/** @return whether we have a next smallest number */
func (bsi *BSTIterator) HasNext() bool {
	return !bsi.stack.isEmpty()
}

func (bsi *BSTIterator) getTreeNodeOfSmallestNumber() *TreeNode{
	return bsi.stack.Pop()
}


func (bsi *BSTIterator) iterate(root *TreeNode) {
	for root != nil {
		bsi.stack.Push(root)
		root = root.Left
	}
}


func main() {

}

/*
	总结
	1. 总体思路就是：通过迭代延迟实现中序遍历
*/
