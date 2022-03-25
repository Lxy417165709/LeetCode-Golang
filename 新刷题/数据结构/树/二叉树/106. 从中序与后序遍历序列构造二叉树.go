package 二叉树

func buildTree(inorder []int, postorder []int) *TreeNode {
	// 1. 空返回。
	if len(postorder) == 0 {
		return nil
	}

	// 2. 获取根节点在中序遍历序列的 index。
	rootNum := postorder[len(postorder)-1]
	rootIndexInInorderSeq := 0
	for index, num := range inorder {
		if num == rootNum {
			rootIndexInInorderSeq = index
			break
		}
	}

	// 3. 递归构建。
	return &TreeNode{
		Val:   rootNum,
		Left:  buildTree(inorder[:rootIndexInInorderSeq], postorder[:rootIndexInInorderSeq]),
		Right: buildTree(inorder[rootIndexInInorderSeq+1:], postorder[rootIndexInInorderSeq:len(postorder)-1]),
	}
}
