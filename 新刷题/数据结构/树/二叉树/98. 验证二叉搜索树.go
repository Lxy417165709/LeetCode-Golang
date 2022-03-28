package 二叉树

import "github.com/Lxy417165709/LeetCode-Golang/新刷题/util/math_util"

// ---------------------------------------- 1. 错误解法(开始) ----------------------------------------

// isValidBST 判断是否是合法的搜索树。 (错误写法，没有考虑根需要大于左子树最大值，小于右子树最小值的条件)
func isValidBST(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left == nil {
		return root.Val < root.Right.Val && isValidBST(root.Right)
	}
	if root.Right == nil {
		return root.Val > root.Left.Val && isValidBST(root.Left)
	}
	return root.Val > root.Left.Val && isValidBST(root.Left) && root.Val < root.Right.Val && isValidBST(root.Right)
}

// ---------------------------------------- 1. 错误解法(结束) ----------------------------------------

// ---------------------------------------- 2. 正确解法(开始) ----------------------------------------

// INF 无穷大。
const INF = 1000000000000

// isValidBST 判断是否是合法的搜索树。 (正确写法)
func isValidBST(root *TreeNode) bool {
	isValid, _, _ := getIsValidBST(root)
	return isValid
}

// getIsValidBST 判断是否是合法的搜索树。
// 返回值:
// 1. 是否是合法的搜索树。
// 2. 树最小值。
// 3. 树最大值。
func getIsValidBST(root *TreeNode) (bool, int, int) {
	// 1. 递归出口。
	if root == nil {
		return true, INF, -INF
	}

	// 2. 向左右子树递归。
	rightIsBST, rightMinValue, rightMaxValue := getIsValidBST(root.Right)
	leftIsBST, leftMinValue, leftMaxValue := getIsValidBST(root.Left)
	return (leftIsBST && root.Val > leftMaxValue) && (rightIsBST && root.Val < rightMinValue),
		math_util.Min(leftMinValue, rightMinValue, root.Val),
		math_util.Max(leftMaxValue, rightMaxValue, root.Val)
}

// ---------------------------------------- 2. 正确解法(结束) ----------------------------------------

// ---------------------------------------- 3. 中序遍历解法(开始) ----------------------------------------

// isValidBST 判断是否是合法的搜索树。 (中序遍历解法)
func isValidBST(root *TreeNode) bool {
	curMaxValue := -INF
	return getIsValidBST(root, &curMaxValue)
}

// isValidBST 判断是否是合法的搜索树。 (中序遍历解法)
func getIsValidBST(root *TreeNode, curMaxValue *int) bool {
	// 1. 递归出口。
	if root == nil {
		return true
	}

	// 2. 左子树判断。
	if !getIsValidBST(root.Left, curMaxValue) {
		return false
	}

	// 3. 判断当前中序遍历最大值。当根节点大于它时，说明中序遍历序列合法，继续判断右子树。
	if root.Val > *curMaxValue {
		*curMaxValue = root.Val
		return getIsValidBST(root.Right, curMaxValue)
	}

	// 4. 返回。
	return false
}

// ---------------------------------------- 3. 中序遍历解法(结束) ----------------------------------------
