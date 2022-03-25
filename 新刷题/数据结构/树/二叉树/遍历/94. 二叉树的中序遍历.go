package 遍历

// inorderTraversal 中序遍历。
func inorderTraversal(root *TreeNode) []int {
	// 1. 空返回。
	if root == nil {
		return nil
	}

	// 2. 初始化。
	hadSonPushIntoStack := make(map[*TreeNode]bool)
	stack := []*TreeNode{root}
	result := make([]int, 0)

	// 3. 中序遍历。
	for len(stack) != 0 {
		// 3.1 出栈。
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top == nil {
			continue
		}

		// 3.2 如果节点的子节点曾经入过栈，此时节点值假如遍历序列。
		if hadSonPushIntoStack[top] {
			result = append(result, top.Val)
			continue
		}

		// 3.3 将节点、节点子节点入栈，并标识。
		stack = append(stack, top.Right, top, top.Left)
		hadSonPushIntoStack[top] = true
	}

	// 4. 返回。
	return result
}
