package 性质判定

// ------------------------ 方法1: 中序遍历 ---------------------------
const INF = 1000000000000

var preNodeValue int

func isValidBST(root *TreeNode) bool {
	preNodeValue = -INF
	return getIsValidBST(root)
}

func getIsValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !getIsValidBST(root.Left) {
		return false
	}
	if preNodeValue >= root.Val {
		return false
	}
	preNodeValue = root.Val
	if !getIsValidBST(root.Right) {
		return false
	}
	return true
}

// ------------------------ 方法2: 先序遍历 ---------------------------
const INF = 1000000000000

func isValidBST(root *TreeNode) bool {
	return getIsValidBST(root, -INF, INF)
}

func getIsValidBST(root *TreeNode, minValue, maxValue int) bool {
	if root == nil {
		return true
	}
	return isInInterval(root.Val, minValue, maxValue) &&
		getIsValidBST(root.Left, minValue, root.Val) &&
		getIsValidBST(root.Right, root.Val, maxValue)
}

func isInInterval(value int, minValue, maxValue int) bool {
	return minValue < value && value < maxValue
}

/*
	题目链接:
	总结:
		1. 这题考查的就是中序遍历，更具体来说就是：二叉搜索树的中序遍历序列是升序的。
		2. 这题和 _面试题 04.05. 合法二叉搜索树_ 是一样的。
		3. 官方有用先序遍历AC这题的。
		3. 官方有用先序遍历AC这题的。
*/
