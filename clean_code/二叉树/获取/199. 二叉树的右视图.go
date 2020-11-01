package 获取

var visibleNodeValues []int

func rightSideView(root *TreeNode) []int {
	visibleNodeValues = make([]int, 0)
	formVisibleNodeValues(root, 0)
	return visibleNodeValues
}

func formVisibleNodeValues(root *TreeNode, nowLay int) {
	if root == nil {
		return
	}
	if nowLay == len(visibleNodeValues) {
		visibleNodeValues = append(visibleNodeValues, root.Val)
	}
	formVisibleNodeValues(root.Right, nowLay+1)
	formVisibleNodeValues(root.Left, nowLay+1)
}

/*
	总结:
		1. 这题考查的是反序的先序遍历。
*/
