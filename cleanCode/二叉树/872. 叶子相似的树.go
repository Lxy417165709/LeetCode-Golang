package 二叉树

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	return areArraysSame(getLeafSequence(root1), getLeafSequence(root2))
}

func getLeafSequence(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	leafSequence := make([]int, 0)
	leafSequence = append(leafSequence, getLeafSequence(root.Left)...)
	leafSequence = append(leafSequence, getLeafSequence(root.Right)...)
	return leafSequence
}

func areArraysSame(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

/*
	题目链接: https://leetcode-cn.com/problems/leaf-similar-trees/submissions/
*/
