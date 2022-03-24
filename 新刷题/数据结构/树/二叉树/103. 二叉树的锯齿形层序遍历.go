package 二叉树

func zigzagLevelOrder(root *TreeNode) [][]int {
	// 1. 获取层序遍历结果。
	levelOrder := levelOrder(root)

	// 2. 翻转偶数层。 (层数从1开始)
	for i := 1; i < len(levelOrder); i += 2 {
		levelOrder[i] = reverseArray(levelOrder[i])
	}
	zigzagLevelOrder := levelOrder

	// 3. 返回。
	return zigzagLevelOrder
}

// reverseArray 翻转数组。
func reverseArray(nums []int) []int {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
	return nums
}
