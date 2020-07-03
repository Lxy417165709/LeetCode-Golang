package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	return getBSTOfSortedArray(nums)
}

func getBSTOfSortedArray(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	middleIndex := len(nums) >> 1
	return &TreeNode{
		nums[middleIndex],
		getBSTOfSortedArray(nums[:middleIndex]),
		getBSTOfSortedArray(nums[middleIndex+1:]),
	}
}

