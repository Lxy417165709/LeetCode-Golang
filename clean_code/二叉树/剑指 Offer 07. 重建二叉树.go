package 二叉树

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	treeBuilder := NewTreeBuilder()
	return treeBuilder.BuildTreeByPreorderAndInorder(preorder, inorder)
}

// ------------ TreeBuilder -----------
type TreeBuilder struct {
}

func NewTreeBuilder() *TreeBuilder {
	return &TreeBuilder{}
}

func (tb *TreeBuilder) BuildTreeByPreorderAndInorder(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	index := getIndexOfSpecificNumInSearchedNums(inorder, rootVal)
	return &TreeNode{
		Val:   rootVal,
		Left:  tb.BuildTreeByPreorderAndInorder(preorder[1:1+index], inorder[:index]),
		Right: tb.BuildTreeByPreorderAndInorder(preorder[1+index:], inorder[index+1:]),
	}
}

func getIndexOfSpecificNumInSearchedNums(searchedNums []int, specificNum int) int {
	for i := 0; i < len(searchedNums); i++ {
		if searchedNums[i] == specificNum {
			return i
		}
	}
	return 0
}
