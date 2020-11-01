package 结构调整

var sumOfGreaterNodes int

func bstToGst(root *TreeNode) *TreeNode {
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
	题目链接: https://leetcode-cn.com/problems/binary-search-tree-to-greater-sum-tree/
*/
