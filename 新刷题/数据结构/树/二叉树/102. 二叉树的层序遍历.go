package 二叉树

// levelOrder 获取二叉树层序遍历序列。
func levelOrder(root *TreeNode) [][]int {
	// 1. 空返回。
	if root == nil {
		return nil
	}

	// 2. BFS获取层序遍历结果。
	result := make([][]int, 0)
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		curLevelNums := make([]int, 0)
		nextQueue := make([]*TreeNode, 0)
		for _, node := range queue {
			curLevelNums = append(curLevelNums, node.Val)
			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}
		queue = nextQueue
		result = append(result, curLevelNums)
	}

	// 3. 返回。
	return result
}
