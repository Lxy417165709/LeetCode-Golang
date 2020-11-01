package 二叉树

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return getMBT(nums)
}

func getMBT(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	indexOfMaxValue := getIndexOfMaxValue(nums)
	return &TreeNode{
		Val:   nums[indexOfMaxValue],
		Left:  getMBT(nums[:indexOfMaxValue]),
		Right: getMBT(nums[indexOfMaxValue+1:]),
	}
}

func getIndexOfMaxValue(nums []int) int {
	indexOfMaxValue := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[indexOfMaxValue] {
			indexOfMaxValue = i
		}
	}
	return indexOfMaxValue
}

/*
	题目链接: https://leetcode-cn.com/problems/maximum-binary-tree/submissions/
	总结:
		1. 可以使用线段树优化 getIndexOfMaxValue 函数。
*/
