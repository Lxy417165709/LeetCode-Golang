package 结构调整

// ---------------------- 递归实现 ----------------------
var sumOfGreaterNodes int

func convertBST(root *TreeNode) *TreeNode {
	sumOfGreaterNodes = 0
	beGreaterTree(root)
	return root
}

func beGreaterTree(root *TreeNode) {
	if root == nil {
		return
	}
	beGreaterTree(root.Right)
	root.Val += sumOfGreaterNodes
	sumOfGreaterNodes = root.Val
	beGreaterTree(root.Left)
}

// ---------------------- 迭代实现 ----------------------
var sumOfGreaterNodes int

func convertBST(root *TreeNode) *TreeNode {
	sumOfGreaterNodes = 0
	beGreaterTree(root)
	return root
}

func beGreaterTree(root *TreeNode) {
	stack := NewTreeNodeStack()
	hadSonNodesPushIntoStack := make(map[*TreeNode]bool)
	stack.Push(root)
	for stack.IsNotEmpty() {
		top := stack.Pop()
		if top == nil {
			continue
		}
		if hadSonNodesPushIntoStack[top] {
			top.Val += sumOfGreaterNodes
			sumOfGreaterNodes = top.Val
		} else {
			stack.Push(top.Left, top, top.Right)
			hadSonNodesPushIntoStack[top] = true
		}
	}
}

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

/*
	题目链接: https://leetcode-cn.com/problems/convert-bst-to-greater-tree/
	总结:
		1. 这题考查的是反向中序遍历。
		2. 这题和 _1038. 从二叉搜索树到更大和树_ 是一样的。
*/
