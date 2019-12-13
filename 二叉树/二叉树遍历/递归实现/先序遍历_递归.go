package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归 + 外部变量 实现先序遍历
var preorderSequence []int
func preorderTraversal(root *TreeNode) []int {
	preorderSequence = make([]int, 0)
	preorderTraversalExec(root)
	return preorderSequence
}

func preorderTraversalExec(root *TreeNode) {
	if root == nil {
		return
	}
	preorderSequence = append(preorderSequence, root.Val)
	preorderTraversalExec(root.Left)
	preorderTraversalExec(root.Right)
}





func main() {

}