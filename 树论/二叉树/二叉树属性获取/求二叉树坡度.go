package main

/*
	给定一个二叉树，计算整个树的坡度。

	一个树的节点的坡度定义即为，该节点左子树的结点之和和右子树结点之和的差的绝对值。空结点的的坡度是0。

	整个树的坡度就是其所有节点的坡度之和。

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 求二叉树坡度		(解法1)
func findTilt(root *TreeNode) int {
	tilt, _ := findTiltExec(root)
	return tilt
}

// 第一个返回值是坡度，第二个是节点和
func findTiltExec(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	leftTilt, leftSum := findTiltExec(root.Left)
	rightTilt, rightSum := findTiltExec(root.Right)
	return leftTilt + rightTilt + abs(leftSum-rightSum), leftSum + rightSum + root.Val
}

// 求二叉树坡度		(解法2)
var tilt int
func findTilt(root *TreeNode) int {
	tilt = 0
	findTiltExec(root)
	return tilt
}

// 返回节点和
func findTiltExec(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftSum := findTiltExec(root.Left)
	rightSum := findTiltExec(root.Right)
	tilt += abs(leftSum - rightSum) // 增加坡度
	return leftSum + rightSum + root.Val
}
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

/*
	题目链接:
		https://leetcode-cn.com/problems/binary-tree-tilt/		二叉树坡度
*/

/*
	总结
	1. 解法1和解法2差不多，只是解法2把坡度保存在一个外部变量，这样就不用通过返回值返回了。
*/
