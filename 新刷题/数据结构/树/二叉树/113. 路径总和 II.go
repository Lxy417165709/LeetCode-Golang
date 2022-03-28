package 二叉树

func pathSum(root *TreeNode, targetSum int) [][]int {
	return getPaths(root, targetSum, nil)
}

// getPaths 获取根节点到叶子节点，和为 targetSum 的路径。
func getPaths(root *TreeNode, targetSum int, curPath []int) [][]int {
	// 1. 空树。
	if root == nil {
		return nil
	}

	// 2. 叶子节点。
	if root.Left == nil && root.Right == nil {
		if targetSum != root.Val {
			return nil
		}
		curPath = append(curPath, root.Val)
		path := make([]int, len(curPath))
		copy(path, curPath)
		return [][]int{path} // 存储时一定要深复制。因为传入的 curPath 地址可能是一样的。 (append可能会使切片扩容，扩容时切片内指针会变，而未扩容时，切片内的指针不变)
	}

	// 3. 非叶子节点。
	result := make([][]int, 0)
	result = append(result, getPaths(root.Left, targetSum-root.Val, append(curPath, root.Val))...)
	result = append(result, getPaths(root.Right, targetSum-root.Val, append(curPath, root.Val))...)

	// 4. 返回。
	return result
}
