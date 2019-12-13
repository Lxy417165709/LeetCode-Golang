package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归 + 外部变量 实现中序遍历
var inorderSequence []int
func inorderTraversal(root *TreeNode) []int {
	inorderSequence = make([]int,0)
	inorderTraversalExec(root)
	return inorderSequence
}

func inorderTraversalExec(root *TreeNode) {
	if root == nil{
		return
	}
	inorderTraversalExec(root.Left)
	inorderSequence = append(inorderSequence,root.Val)	// 与先序遍历相比，就这语句调整了下位置。
	inorderTraversalExec(root.Right)
}





func main() {

}