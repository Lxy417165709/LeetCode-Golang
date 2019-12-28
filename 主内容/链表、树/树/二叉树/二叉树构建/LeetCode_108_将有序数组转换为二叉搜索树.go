package main

func sortedArrayToBST(nums []int) *TreeNode {
	return sortedArrayToBSTExec(nums)
}

func sortedArrayToBSTExec(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) >> 1
	root := &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBSTExec(nums[:mid]),
		Right: sortedArrayToBSTExec(nums[mid+1:]),
	}
	return root
}

/*
	题目链接:
		https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/submissions/		将有序数组转换为二叉搜索树
*/

/*
	总结
	1. 这题其实就是通过中序遍历序列构建二叉搜索树，但它多了一个条件，这颗二叉树要平衡，即左右高度差不超过1。
	   这很简单，我们每次都取序列的中点作为根节点，之后再递归构建左右子树就可以了。
	2. 这题还可以使用迭代实现，但是代码比递归的难一些。
*/