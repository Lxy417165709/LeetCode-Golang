package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return constructMaximumBinaryTreeExec(nums)
}

func constructMaximumBinaryTreeExec(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	rootIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[rootIndex] {
			rootIndex = i
		}
	}
	root := &TreeNode{
		Val:   nums[rootIndex],
		Left:  constructMaximumBinaryTreeExec(nums[:rootIndex]),
		Right: constructMaximumBinaryTreeExec(nums[rootIndex+1:]),
	}
	return root
}

/*
	题目链接:
		https://leetcode-cn.com/problems/maximum-binary-tree/		最大二叉树
*/

/*
	总结
	1. 这题只需先将序列遍历一次，找出最大值的索引，之后将序列划分为两部分，再进行递归建左右子树就可以了。
*/
