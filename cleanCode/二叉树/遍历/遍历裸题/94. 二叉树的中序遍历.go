package 遍历裸题

func inorderTraversal(root *TreeNode) []int {
	sequence := make([]int, 0)
	stack := NewTreeNodeStack()
	hadSonNodesPushIntoStack := make(map[*TreeNode]bool)
	stack.Push(root)
	for stack.IsNotEmpty() {
		top := stack.Pop()
		if top == nil {
			continue
		}
		if hadSonNodesPushIntoStack[top] {
			sequence = append(sequence, top.Val)
		} else {
			stack.Push(top.Right, top, top.Left)
			hadSonNodesPushIntoStack[top] = true
		}
	}
	return sequence
}

// ----------------------------- TreeNodeStack -----------------------------
type TreeNodeStack struct {
	data []*TreeNode
}

func NewTreeNodeStack() *TreeNodeStack {
	return &TreeNodeStack{}
}

func (tns *TreeNodeStack) Push(nodes ...*TreeNode) {
	tns.data = append(tns.data, nodes...)
}

func (tns *TreeNodeStack) IsNotEmpty() bool {
	return tns.getSize() != 0
}

func (tns *TreeNodeStack) Pop() *TreeNode {
	topNode := tns.data[tns.getSize()-1]
	tns.data = tns.data[:tns.getSize()-1]
	return topNode
}

func (tns *TreeNodeStack) getSize() int {
	return len(tns.data)
}

// ----------------------------- TreeNode -----------------------------
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	题目链接: https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
	总结:
		1. 对本题代码重构后，消除了变量命名上的错误，使得代码更加清晰易于理解。
*/
